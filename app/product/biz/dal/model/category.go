package model

import (
	"context"
	// "fmt"
	"time"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
	"github.com/whlxbd/gomall/common/utils/pool"
	"gorm.io/gorm"
)

type Category struct {
	Base
	Name        string `json:"name" gorm:"uniqueIndex;type:varchar(50);not null"` // 分类名称(唯一索引)
	Description string `json:"description"`                                       // 分类描述

	Products []Product `gorm:"many2many:product_category;"` // 分类下的商品
}

func (c Category) TableName() string {
	return "category"
}

type CategoryQuery struct {
	ctx context.Context
	db  *gorm.DB
}

type CachedCategoryQuery struct {
	categoryQuery *CategoryQuery
	cacheClient   *redis.Client
	prefix        string
}

// NewCategoryQuery 创建一个新的CategoryQuery
func NewCategoryQuery(ctx context.Context, db *gorm.DB) *CategoryQuery {
	return &CategoryQuery{
		ctx: ctx,
		db:  db,
	}
}

// NewCachedCategoryQuery 创建一个新的CachedCategoryQuery
func NewCachedCategoryQuery(cq *CategoryQuery, cacheClient *redis.Client) *CachedCategoryQuery {
	return &CachedCategoryQuery{
		categoryQuery: cq,
		cacheClient:   cacheClient,
		prefix:        "gomall_category_",
	}
}

// 在数据库中检查分类是否存在
func (cq CategoryQuery) IsExistByName(name string) (uint32, error) {
	var id uint32
	err := cq.db.WithContext(cq.ctx).Model(&Category{}).Where("name = ?", name).Select("id").Find(&id).Error
	if err != nil {
		return 0, err
	}

	return id, nil
}

// 设置缓存
func (ccq CachedCategoryQuery) SetCache(name string, id uint32) error {
	key := ccq.prefix + name
	return ccq.cacheClient.Set(ccq.categoryQuery.ctx, key, id, 24*time.Hour).Err()
}

// 检查是否存在
func (ccq CachedCategoryQuery) getFromCache(name string) (uint32, error) {
	key := ccq.prefix + name

	id, err := ccq.cacheClient.Get(ccq.categoryQuery.ctx, key).Uint64()
	if err != nil {
		return 0, err
	}

	return uint32(id), nil
}

func (ccq CachedCategoryQuery) IsExistByName(name string) (uint32, error) {
	// 1. 先查缓存
	id, err := ccq.getFromCache(name)
	if err == nil && id > 0 {
		_ = pool.Submit(func() {
			if err := ccq.SetCache(name, id); err != nil {
				klog.Error("设置缓存失败", err)
			}
		})
		// fmt.Println("缓存中存在")
		return id, nil // 缓存中存在
	}
	if err != redis.Nil && err != nil {
		return 0, err // 发生错误
	}

	// 2. 缓存未命中，查数据库
	id, err = ccq.categoryQuery.IsExistByName(name)
	if err != nil {
		return 0, err
	}

	// 3. 设置缓存
	if id > 0 {
		_ = pool.Submit(func() {
			if err := ccq.SetCache(name, id); err != nil {
				klog.Error("设置缓存失败", err)
			}
		})
		// fmt.Println("数据库中存在")
	} else if id == 0 {
		// fmt.Println("数据库中不存在")
		return 0, nil
	}

	return id, nil
}

func GetByCategoryName(db *gorm.DB, ctx *context.Context, categoryId uint32, page int32, pageSize int64) (*[]Product, error) {
	var products []Product
	query := db.WithContext(*ctx).Model(&Product{})

    if categoryId > 0 {
		query = query.
		Model(&Product{}).
		Joins("JOIN product_category ON product.id = product_category.product_id").
		Where("product_category.category_id = ?", categoryId)
	}

	// 分页查询商品
    if err := query.
        Preload("Categories").
        Limit(int(pageSize)).
        Offset(int((page - 1) * int32(pageSize))).
        Find(&products).Error; err != nil {
		klog.Errorf("查询商品失败: %v", err)
		return nil, kerrors.NewBizStatusError(400, "查询商品失败")
	}

	return &products, nil
}

// 流式查询商品列表
func StreamProductsByCategory(db *gorm.DB, ctx *context.Context, categoryId uint32, handleFunc func(*Product) error) error {
    // 构建基础查询
    query := db.WithContext(*ctx).
        Model(&Product{}).
        Preload("Categories")

    // ID为0时查询所有商品
    if categoryId > 0 {
        query = query.
            Joins("JOIN product_category ON products.id = product_category.product_id").
            Where("product_category.category_id = ?", categoryId)
    }

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