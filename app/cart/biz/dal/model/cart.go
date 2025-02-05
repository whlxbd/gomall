package model

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// redis策略 Cache Aside Pattern

type Cart struct {
	Base
	UserID    uint32 `json:"user_id" gorm:"index:idx_user_product;not null;comment:用户ID"`
	ProductID uint32 `json:"product_id" gorm:"index:idx_user_product;not null;comment:商品ID"`
	Quantity  uint32 `json:"quantity" gorm:"not null;default:1;check:quantity > 0;comment:商品数量"`
	Selected  bool   `json:"selected" gorm:"default:true;comment:是否选中"`
	Status    bool   `json:"status" gorm:"default:true;comment:状态 true:正常 false:失效"`
}

const (
	CartCachePrefix = "gomall_cart:userid:"
	CartExpiration  = time.Minute * 10
)

func (c Cart) TableName() string {
	return "cart"
}

type CartMessage struct {
	Operation string  `json:"operation"`
	UserID    uint32  `json:"user_id"`
	CartList  []*Cart `json:"cart"`
}

// 从缓存获取购物车信息，若没有则从数据库获取
func GetCartByUserId(db *gorm.DB, cacheClient *redis.Client, ctx context.Context, userId uint32) (cartList []*Cart, err error) {
	cartList, err = CachedGetCartByUserId(cacheClient, ctx, userId)
	if err == nil {
		return
	}
	if err != redis.Nil {
		klog.Error("获取购物车缓存失败", err)
	}

	// 从数据库获取
	err = db.Debug().WithContext(ctx).Model(&Cart{}).Where("user_id = ?", userId).Find(&cartList).Error

	return
}

// 从缓存获取购物车信息
func CachedGetCartByUserId(cacheClient *redis.Client, ctx context.Context, userId uint32) (cartList []*Cart, err error) {
	key := CartCachePrefix + strconv.Itoa(int(userId))

	pipe := cacheClient.Pipeline()
	getCmd := pipe.Get(ctx, key)
	pipe.Expire(ctx, key, CartExpiration) // 更新过期时间

	_, err = pipe.Exec(ctx)
	if err != nil {
		return nil, err
	}

	data, err := getCmd.Result()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(data), &cartList)
	return
}

func AddCart(cart *Cart, db *gorm.DB, ctx context.Context, cacheClient *redis.Client) error {
	return nil
}
