package model

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
	"github.com/whlxbd/gomall/common/utils/pool"
	"gorm.io/gorm"
)

// redis策略 Read/Write Through Pattern

// 商品状态: 1:上架 2:下架 3:删除
type ProductStatus int32

const (
	ProductStatusOnSale  ProductStatus = 1
	ProductStatusOffSale ProductStatus = 2
	ProductStatusDeleted ProductStatus = 3
)

func (s ProductStatus) IsValid() error {
	if s < ProductStatusOnSale || s > ProductStatusDeleted {
		return fmt.Errorf("invalid ProductStatus: %d", s)

	}
	return nil

}

type Product struct {
	Base
	Name        string  `json:"name"`        // 商品名称
	Description string  `json:"description"` // 商品描述
	Picture     string  `json:"picture"`     // 商品图片
	Price       float32 `json:"price"`       // 商品价格

    Stock     int32 `json:"stock" gorm:"check:stock >= 0;not null;default:0"`   // 库存数量
	SoldCount int32 `json:"sold_count"` // 销售数量

	Status      ProductStatus `json:"status"`       // 商品状态(1:上架 2:下架 3:删除)
	IsHot       bool          `json:"is_hot"`       // 是否热销
	IsNew       bool          `json:"is_new"`       // 是否新品
	IsRecommend bool          `json:"is_recommend"` // 是否推荐

	Categories []Category `gorm:"many2many:product_category;"`
}

type ProductStockItem struct {
    ProductId uint32
	SoldcountChange int32 // 销售数量变化，负数表示退货
    StockChange  int32 // 数量变化，负数表示扣减库存
}

type ProductQuery struct {
	ctx context.Context // 上下文
	db  *gorm.DB        // 数据库连接
}

type CachedProductQuery struct {
	productQuery *ProductQuery // 商品查询
	cacheClient  *redis.Client // 缓存客户端
	prefix       string        // 缓存前缀
}

func (p Product) TableName() string {
	return "product"
}

func (p *Product) BeforeMigrate(tx *gorm.DB) error {
    return tx.Exec(`
        ALTER TABLE products
        ADD CONSTRAINT check_stock_non_negative CHECK (stock >= 0),
        ADD CONSTRAINT check_sold_count_non_negative CHECK (sold_count >= 0),
        ADD CONSTRAINT check_price_non_negative CHECK (price >= 0)
    `).Error
}

// 从数据库获取商品
func (p ProductQuery) GetById(productid uint32) (*Product, error) {
	product := new(Product)
	err := p.db.WithContext(p.ctx).Where(&Product{Base: Base{ID: productid}}).First(product).Error
	return product, err
}

// 批量获取商品
func (p ProductQuery) GetByIds(productIds []uint32) ([]*Product, error) {
	var products []*Product
	err := p.db.WithContext(p.ctx).
		Where("id IN ?", productIds).
		Preload("Categories").
		Find(&products).Error
	return products, err
}

// 尝试从缓存获取商品，如果缓存不存在则从数据库获取
func (p CachedProductQuery) GetById(productid uint32) (*Product, error) {
	key := p.prefix + strconv.FormatUint(uint64(productid), 10)

	// 尝试从缓存获取
	product, err := p.getFromCache(key)
	if err == nil {
		// 更新缓存
		_ = pool.Submit(func() {
			if err := p.setCache(key, product); err != nil {
				klog.Error("缓存存在但设置缓存失败", err)
			}
		})
		return product, nil
	}
	if err != redis.Nil {
		return &Product{}, kerrors.NewBizStatusError(400, err.Error())
	}

	// 从数据库获取
	product, err = p.productQuery.GetById(productid)
	if err != nil {
		return &Product{}, kerrors.NewBizStatusError(400, err.Error())

	}

	// 更新缓存
	_ = pool.Submit(func() {
		if err := p.setCache(key, product); err != nil {
			klog.Error("缓存不存在，设置缓存失败", err)
		}
	})

	return product, nil
}

// 批量获取商品（带缓存）
func (p CachedProductQuery) GetByIds(productIds []uint32) ([]*Product, error) {
	if len(productIds) == 1 {
		product, err := p.GetById(productIds[0])
		if err != nil {
			return nil, err
		}
		return []*Product{product}, nil
	}

	products := make([]*Product, 0, len(productIds))
	missingIds := make([]uint32, 0)

	// 1. 尝试从缓存获取
	for _, id := range productIds {
		if id == 0 {
			continue
		}
		key := p.prefix + strconv.FormatUint(uint64(id), 10)
		product, err := p.getFromCache(key)
		if err == nil {
			products = append(products, product)
			_ = pool.Submit(func() {
				if err := p.setCache(key, product); err != nil {
					klog.Error("设置缓存失败", err)
				}
			})
		} else if err == redis.Nil {
			missingIds = append(missingIds, id)
		} else {
			return nil, kerrors.NewBizStatusError(400, err.Error())
		}
	}

	// 2. 未命中的从数据库获取
	if len(missingIds) > 0 {
		dbProducts, err := p.productQuery.GetByIds(missingIds)
		if err != nil {
			return nil, err
		}
		products = append(products, dbProducts...)

		// 3. 异步写入缓存
		for _, product := range dbProducts {
			product := product // 避免闭包问题
			_ = pool.Submit(func() {
				key := p.prefix + strconv.FormatUint(uint64(product.ID), 10)
				if err := p.setCache(key, product); err != nil {
					klog.Error("设置缓存失败", err)
				}
			})
		}
	}

	return products, nil
}

// 从缓存中获取商品
func (p CachedProductQuery) getFromCache(key string) (*Product, error) {
	product := new(Product)
	val, err := p.cacheClient.Get(p.productQuery.ctx, key).Result()
	if err != nil {
		return &Product{}, err
	}

	if err := json.Unmarshal([]byte(val), &product); err != nil {
		return &Product{}, err
	}
	return product, nil
}

// 设置商品缓存
func (p CachedProductQuery) setCache(key string, product *Product) error {
	encoded, err := json.Marshal(product)
	if err != nil {
		return err
	}

	return p.cacheClient.Set(p.productQuery.ctx, key, encoded, time.Hour).Err()
}

// 创建一个商品查询
func NewProductQuery(ctx context.Context, db *gorm.DB) *ProductQuery {
	return &ProductQuery{ctx: ctx, db: db}
}

// 创建一个带缓存的商品查询
func NewCachedProductQuery(productQuery *ProductQuery, cacheClient *redis.Client) *CachedProductQuery {
	return &CachedProductQuery{productQuery: productQuery, cacheClient: cacheClient, prefix: "gomall_product:productid:"}
}

// 通过商品ID获取商品
func GetById(ctx context.Context, db *gorm.DB, productid uint32) (product *Product, err error) {
	err = db.WithContext(ctx).Where(&Product{Base: Base{ID: productid}}).First(product).Error
	return
}

// 通过商品名称或描述搜索商品
func SearchProduct(db *gorm.DB, ctx *context.Context, q string, page int32, pageSize int64) (products []*Product, err error) {
	query := db.WithContext(*ctx).Model(&Product{}).
		Select("id, name, description, price, stock").
		Where("name like ? or description like ?", "%"+q+"%", "%"+q+"%")

	// 分页查询
	if err = query.Preload("Categories").
		Limit(int(pageSize)).
		Offset(int((page - 1) * int32(pageSize))).
		Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

// 流式搜索商品
func StreamSearchProduct(db *gorm.DB, ctx *context.Context, q string, handleFunc func(*Product) error) error {
	query := db.WithContext(*ctx).
		Model(&Product{}).
		Preload("Categories").
		Where("name like ? or description like ?", "%"+q+"%", "%"+q+"%")

	rows, err := query.Rows()
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var product Product
		if err := db.ScanRows(rows, &product); err != nil {
			return err
		}
		if err := handleFunc(&product); err != nil {
			return err
		}
	}

	return nil
}

// 创建商品
func CreateProduct(db *gorm.DB, cacheClient *redis.Client, ctx context.Context, product *Product) error {
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(product).Error; err != nil {
			klog.Errorf("CreateProduct error: %v", err)
			return kerrors.NewBizStatusError(400, "创建商品失败")
		}

		// 创建缓存
		_ = pool.Submit(func() {
			p := NewCachedProductQuery(NewProductQuery(ctx, db), cacheClient)
			key := p.prefix + strconv.FormatUint(uint64(product.ID), 10)
			if err := p.setCache(key, product); err != nil {
				klog.Error("设置缓存失败", err)
			}
		})

		return nil
	})
}

// 更新商品
func EditProduct(db *gorm.DB, cacheClient *redis.Client, ctx context.Context, product *Product) error {
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(product).Updates(product).Error; err != nil {
			klog.Errorf("EditProduct error: %v", err)
			return kerrors.NewBizStatusError(400, "更新商品失败")
		}

		// 更新缓存
		_ = pool.Submit(func() {
			p := NewCachedProductQuery(NewProductQuery(ctx, db), cacheClient)
			key := p.prefix + strconv.FormatUint(uint64(product.ID), 10)
			if err := p.setCache(key, product); err != nil {
				klog.Error("设置缓存失败", err)
				// 删除缓存
				if err := p.cacheClient.Del(p.productQuery.ctx, key).Err(); err != nil {
					klog.Error("删除缓存失败", err)
				}
			}
		})

		return nil
	})
}

// 删除商品
func DeleteProduct(db *gorm.DB, cacheClient *redis.Client, ctx context.Context, productId uint32) error {
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Product{Base: Base{ID: productId}}).Association("Categories").Clear(); err != nil {
			klog.Errorf("删除商品分类关联失败: %v", err)
			return kerrors.NewBizStatusError(400, "删除商品分类关联失败")
		}

		if err := tx.Delete(&Product{Base: Base{ID: productId}}).Error; err != nil {
			klog.Errorf("DeleteProduct error: %v", err)
			return kerrors.NewBizStatusError(400, "更新商品失败")
		}

		// 删除缓存
		_ = pool.Submit(func() {
			p := NewCachedProductQuery(NewProductQuery(ctx, db), cacheClient)
			key := p.prefix + strconv.FormatUint(uint64(productId), 10)
			if err := p.cacheClient.Del(p.productQuery.ctx, key).Err(); err != nil {
				klog.Error("删除缓存失败", err)
			}
		})

		return nil
	})
}

func UpdateBatchProduct(db *gorm.DB, cacheClient *redis.Client, ctx *context.Context, items []*ProductStockItem, isStock bool) error {
	return db.WithContext(*ctx).Transaction(func(tx *gorm.DB) error {
		ids := make([]string, 0, len(items))
		var stockCase strings.Builder
		var cnt int

		stockCase.WriteString("CASE id")

		if isStock {
			for _, item := range items {
				ids = append(ids, fmt.Sprintf("%d", item.ProductId))

				if item.StockChange == 0 {
					cnt += 1
					continue
				} else if item.StockChange < 0 {
					stockCase.WriteString(fmt.Sprintf(" WHEN %d THEN stock %d", item.ProductId, item.StockChange))
				} else {
					stockCase.WriteString(fmt.Sprintf(" WHEN %d THEN stock + %d", item.ProductId, item.StockChange))
				}
			}
		} else {
			for _, item := range items {
				ids = append(ids, fmt.Sprintf("%d", item.ProductId))

				if item.SoldcountChange == 0 {
					cnt += 1
					continue
				} else if item.SoldcountChange < 0 {
					stockCase.WriteString(fmt.Sprintf(" WHEN %d THEN sold_count %d", item.ProductId, item.SoldcountChange))
				} else {
					stockCase.WriteString(fmt.Sprintf(" WHEN %d THEN sold_count + %d", item.ProductId, item.SoldcountChange))
				}
			}
		}

		if cnt == len(items) {
			return nil
		}
		stockCase.WriteString(" END")

		idList := strings.Join(ids, ",")
		var err error
		if isStock {
			err = tx.Exec(fmt.Sprintf("UPDATE product SET stock = %s WHERE id IN (%s)", stockCase.String(), idList)).Error
		} else {
			err = tx.Exec(fmt.Sprintf("UPDATE product SET sold_count = %s WHERE id IN (%s)", stockCase.String(), idList)).Error
		}

		if err != nil {
			_ = pool.Submit(func() {
				klog.Errorf("UpdateBatchProduct error: %v", err)
				
			})
			if isStock {
				return kerrors.NewBizStatusError(40010, "商品库存不足")
			} else {
				return kerrors.NewBizStatusError(40011, "销量修改失败")
			}
		}

		// 删除缓存
		_ = pool.Submit(func() {
			for _, item := range items {
				key := "gomall_product:productid:" + strconv.FormatUint(uint64(item.ProductId), 10)
				if err := cacheClient.Del(*ctx, key).Err(); err != nil {
					klog.Error("删除缓存失败", err)
				}
			}
		})

		return err
	})
}