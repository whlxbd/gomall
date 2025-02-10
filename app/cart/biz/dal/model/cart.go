package model

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/whlxbd/gomall/common/utils/pool"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// redis策略 Cache Aside Pattern

type Cart struct {
	Base
	UserID    uint32 `json:"user_id" gorm:"index:idx_user_product;not null;comment:用户ID"`
	ProductID uint32 `json:"product_id" gorm:"index:idx_user_product;not null;comment:商品ID"`
	Quantity  int32 `json:"quantity" gorm:"not null;default:1;check:quantity > 0;comment:商品数量"`
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

// 设置购物车缓存
func SetCartCache(cacheClient *redis.Client, ctx context.Context, userId uint32, cartList []*Cart)  {
	key := CartCachePrefix + strconv.Itoa(int(userId))

	data, err := json.Marshal(cartList)
	if err != nil {
		klog.Errorf("购物车缓存序列化失败: %v", err)
	}

	_, err = cacheClient.Set(ctx, key, data, CartExpiration).Result()
	if err != nil {
		klog.Errorf("购物车缓存设置失败: %v", err)
	}
}

// 从缓存获取购物车信息，若没有则从数据库获取
func GetCartByUserId(userId uint32, db *gorm.DB, ctx context.Context, cacheClient *redis.Client) (cartList []*Cart, err error) {
	cartList, err = CachedGetCartByUserId(cacheClient, ctx, userId)
	if err == nil {
		return
	}
	if err != redis.Nil {
		klog.Error("获取购物车缓存失败", err)
	}

	// 从数据库获取
	err = db.Debug().WithContext(ctx).Model(&Cart{}).Where("user_id = ?", userId).Find(&cartList).Error
	if err != nil {
		klog.Errorf("获取购物车失败: %v", err)
		return nil, kerrors.NewBizStatusError(500, "获取购物车失败")
	}

	// 设置缓存
	_ = pool.Submit(func() {
		SetCartCache(cacheClient, ctx, userId, cartList)
	})
	
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

func AddCart(c *Cart, db *gorm.DB, ctx context.Context, cacheClient *redis.Client) error {
	var find Cart
	err := db.WithContext(ctx).Model(&Cart{}).Where(&Cart{UserID: c.UserID, ProductID: c.ProductID}).First(&find).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if find.ID != 0 {
		err = db.WithContext(ctx).Model(&Cart{}).Where(&Cart{UserID: c.UserID, ProductID: c.ProductID}).UpdateColumn("Quantity", gorm.Expr("Quantity+?", c.Quantity)).Error
	} else {
		err = db.WithContext(ctx).Model(&Cart{}).Create(c).Error
	}

	if err != nil {
		klog.Errorf("添加购物车失败: %v", err)
		return kerrors.NewBizStatusError(500, "添加购物车失败")
	}

	// 删除缓存
	_ = pool.Submit(func() {
		key := CartCachePrefix + strconv.Itoa(int(c.UserID))

		_, err = cacheClient.Del(ctx, key).Result()
	})

	return nil
}

func EmptyCart(userId uint32, db *gorm.DB, ctx context.Context, cacheClient *redis.Client) error {
	err := db.WithContext(ctx).Model(&Cart{}).Where("user_id = ?", userId).Delete(&Cart{}).Error
	if err != nil {
		klog.Errorf("清空购物车失败: %v", err)
		return kerrors.NewBizStatusError(500, "清空购物车失败")
	}

	// 删除缓存
	_ = pool.Submit(func() {
		key := CartCachePrefix + strconv.Itoa(int(userId))

		_, err = cacheClient.Del(ctx, key).Result()
	})
	return nil
}