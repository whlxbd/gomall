package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/whlxbd/gomall/common/utils/pool"
	"gorm.io/gorm"
	"github.com/cloudwego/kitex/pkg/klog"
)

type Category struct {
    Base
    Name        string    `json:"name" gorm:"uniqueIndex;type:varchar(50);not null"` // 分类名称(唯一索引)
    Description string    `json:"description"`                                        // 分类描述
    
	Products    []Product `gorm:"many2many:product_category;"`                       // 分类下的商品
}

func (c Category) TableName () string {
	return "category"
}

type CategoryQuery struct {
	ctx context.Context
	db *gorm.DB
}

type CachedCategoryQuery struct {
	categoryQuery *CategoryQuery
	cacheClient *redis.Client
	prefix string
}

// NewCategoryQuery 创建一个新的CategoryQuery
func NewCategoryQuery(ctx context.Context, db *gorm.DB) *CategoryQuery {
	return &CategoryQuery{
		ctx: ctx,
		db: db,
	}
}

// NewCachedCategoryQuery 创建一个新的CachedCategoryQuery
func NewCachedCategoryQuery(ctx context.Context, db *gorm.DB, cacheClient *redis.Client) *CachedCategoryQuery {
	return &CachedCategoryQuery{
		categoryQuery: NewCategoryQuery(ctx, db),
		cacheClient: cacheClient,
		prefix: "gomall_category_",
	}
}

// 在数据库中检查分类是否存在
func (cq CategoryQuery) IsExistByName (name string) (bool, error) {
	var count int64
	if err := cq.db.Model(&Category{}).Where("name = ?", name).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil	
}

// 设置缓存
func (ccq CachedCategoryQuery) setCache(name string) error {
    key := ccq.prefix + name
    return ccq.cacheClient.Set(ccq.categoryQuery.ctx, key, "1", time.Hour).Err()
}

// 检查是否存在
func (ccq CachedCategoryQuery) getFromCache (name string) (bool, error) {
	key := ccq.prefix + name
    _, err := ccq.cacheClient.Get(ccq.categoryQuery.ctx, key).Result()
    if err == redis.Nil {
        return false, redis.Nil
    }
    if err != nil {
        return false, err
    }
    return true, nil
}

func (ccq CachedCategoryQuery) IsExistByName(name string) (bool, error) {
    // 1. 先查缓存
    exists, err := ccq.getFromCache(name)
    if err == nil {
        return true, nil  // 缓存中存在
    }
    if err != redis.Nil {
        return false, err // 发生错误
    }

    // 2. 缓存未命中，查数据库
    exists, err = ccq.categoryQuery.IsExistByName(name)
    if err != nil {
        return false, err
    }

    // 3. 如果数据库中存在，则设置缓存
    if exists {
        _ = pool.Submit(func() {
            if err := ccq.setCache(name); err != nil {
                klog.Error("设置缓存失败", err)
            }
        })
    }

    return exists, nil
}