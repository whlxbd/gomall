package model

import (
	"context"
	"encoding/json"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Product struct {
	Base
	Name        string  `json:"name"`        // 商品名称
	Description string  `json:"description"` // 商品描述
	Picture     string  `json:"picture"`     // 商品图片
	Price       float32 `json:"price"`       // 商品价格

	Stock     	int32 `json:"stock"`      	// 库存数量
	SoldCount 	int32 `json:"sold_count"` 	// 销售数量

	Status      int8 `json:"status"`       // 商品状态(1:上架 2:下架 3:删除)
	IsHot       bool `json:"is_hot"`       // 是否热销
	IsNew       bool `json:"is_new"`       // 是否新品
	IsRecommend bool `json:"is_recommend"` // 是否推荐

	Categories []Category `gorm:"many2many:product_category;"`
}

type ProductQuery struct {
	ctx context.Context // 上下文
	db  *gorm.DB        // 数据库连接
}

type CachedProductQuery struct {
	productQuery ProductQuery  // 商品查询
	cacheClient  *redis.Client // 缓存客户端
	prefix       string        // 缓存前缀
}

func (p Product) TableName() string {
	return "product"
}

// 从数据库获取商品
func (p ProductQuery) GetById(productid int32) (product Product, err error) {
	err = p.db.WithContext(p.ctx).Where(&Product{Base: Base{ID: productid}}).First(&product).Error
	return
}

// 尝试从缓存获取商品，如果缓存不存在则从数据库获取
func (p CachedProductQuery) GetById(productid int32) (Product, error) {
	key := p.prefix + "_product:_" + string(productid)

	// 尝试从缓存获取
	product, err := p.getFromCache(key)
	if err == nil {
		// 更新缓存
		if err := p.setCache(key, product); err != nil {
			klog.Error("设置缓存失败", err)
		}
		return product, nil
	}
	if err != redis.Nil {
		return Product{}, err
	}

	// 从数据库获取
	product, err = p.productQuery.GetById(productid)
	if err != nil {
		return Product{}, err
	}

	// 更新缓存
	if err := p.setCache(key, product); err != nil {
		klog.Error("设置缓存失败", err)
	}

	return product, nil
}

// 从缓存中获取商品
func (p CachedProductQuery) getFromCache(key string) (Product, error) {
	var product Product
	val, err := p.cacheClient.Get(p.productQuery.ctx, key).Result()
	if err != nil {
		return Product{}, err
	}

	if err := json.Unmarshal([]byte(val), &product); err != nil {
		return Product{}, err
	}
	return product, nil
}

// 设置商品缓存
func (p CachedProductQuery) setCache(key string, product Product) error {
	encoded, err := json.Marshal(product)
	if err != nil {
		return err
	}

	return p.cacheClient.Set(p.productQuery.ctx, key, encoded, time.Hour).Err()
}

// 创建一个商品查询
func NewProductQuery(ctx context.Context, db *gorm.DB) ProductQuery {
	return ProductQuery{ctx: ctx, db: db}
}

// 创建一个带缓存的商品查询
func NewCachedProductQuery(productQuery ProductQuery, cacheClient *redis.Client, prefix string) CachedProductQuery {
	return CachedProductQuery{productQuery: productQuery, cacheClient: cacheClient, prefix: "gomall"}
}

// 通过商品ID获取商品
func GetById(ctx context.Context, db *gorm.DB, productid int32) (product Product, err error) {
	err = db.WithContext(ctx).Where(&Product{Base: Base{ID: productid}}).First(&product).Error
	return
}

// 通过商品名称或描述搜索商品
func SearchProduct(db *gorm.DB, ctx context.Context, q string) (product []*Product, err error) {
	err = db.WithContext(ctx).Model(&Product{}).Find(&product, "name like ? or description like ?", "%"+q+"%", "%"+q+"%").Error
	return product, err
}
