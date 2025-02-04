package model

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	product "github.com/whlxbd/gomall/app/product/biz/dal/model"
	"gorm.io/gorm"
)

type Cart struct {
	Base
	UserID    uint32 `json:"user_id" gorm:"index:idx_user_product;not null;comment:用户ID"`
	ProductID uint32 `json:"product_id" gorm:"index:idx_user_product;not null;comment:商品ID"`
	Quantity  uint32 `json:"quantity" gorm:"not null;default:1;check:quantity > 0;comment:商品数量"`
	Selected  bool   `json:"selected" gorm:"default:true;comment:是否选中"`
	Status    int8   `json:"status" gorm:"default:1;comment:状态 1:正常 2:失效"`

	// 关联商品信息(预加载)
	Product *product.Product `json:"product" gorm:"foreignKey:ProductID"`
}

func (c Cart) TableName() string {
	return "cart"
}

type CartQuery struct {
	ctx context.Context // 上下文
	db  *gorm.DB        // 数据库连接
}

type CachedCartQuery struct {
	cartQuery   *CartQuery    // 购物车查询
	cacheClient *redis.Client // 缓存客户端
	prefix      string        // 缓存前缀
}

func NewCartQuery(ctx context.Context, db *gorm.DB) *CartQuery {
	return &CartQuery{ctx: ctx, db: db}
}

func NewCachedCartQuery(cartQuery *CartQuery, cacheClient *redis.Client) *CachedCartQuery {
	return &CachedCartQuery{cartQuery: cartQuery, cacheClient: cacheClient, prefix: "gomall_cart_"}
}

// 根据用户ID从数据库获取购物车
func GetByUserID(userid uint32, db *gorm.DB, ctx context.Context) (*[]Cart, error) {
	cart := new([]Cart)
	err := db.WithContext(ctx).Where(&Cart{UserID: userid}).Find(cart).Error
	return cart, err
}

func GetByCartId(cartid uint32, db *gorm.DB, ctx context.Context) (*Cart, error) {
	cart := new(Cart)
	err := db.WithContext(ctx).Where(&Cart{Base: Base{ID: cartid}}).First(cart).Error
	return cart, err
}

// 从数据库获取购物车
func (c CartQuery) GetByCartId(cartid uint32) (*Cart, error) {
	cart := new(Cart)
	err := c.db.WithContext(c.ctx).Where(&Cart{Base: Base{ID: cartid}}).First(cart).Error
	return cart, err
}

// 从数据库获取用户购物车
func (c CartQuery) GetByUserID(userid uint32) (*[]Cart, error) {
	carts := new([]Cart)
	err := c.db.WithContext(c.ctx).Where(&Cart{UserID: userid}).Find(carts).Error
	return carts, err
}

func (c CachedCartQuery) GetFromCacheByCartId(cartid uint32) (*Cart, error) {
	cart := new(Cart)
	key := c.prefix + "cartid_" + strconv.FormatUint(uint64(cartid), 10)
	err := c.cacheClient.Get(c.cartQuery.ctx, key).Scan(cart)

	if err != nil {
		return nil, err
	}
	return cart, nil
}

func (c CachedCartQuery) GetFromCacheByUserID(userid uint32) (*[]Cart, error) {
	carts := new([]Cart)
	key := c.prefix + "userid_" + strconv.FormatUint(uint64(userid), 10)
	err := c.cacheClient.Get(c.cartQuery.ctx, key).Scan(carts)

	if err != nil {
		return nil, err
	}
	return carts, nil
}

func (c CachedCartQuery) SetCacheByCartId(cartid uint32, cart *Cart) error {
	encoded, err := json.Marshal(cart)
	if err != nil {
		return err
	}

	key := c.prefix + "cartid_" + strconv.FormatUint(uint64(cartid), 10)
	return c.cacheClient.Set(c.cartQuery.ctx, key, encoded, time.Hour).Err()
}

func (c CachedCartQuery) SetCacheByUserID(userid uint32, carts *[]Cart) error {
	encoded, err := json.Marshal(carts)
	if err != nil {
		return err
	}

	key := c.prefix + "userid_" + strconv.FormatUint(uint64(userid), 10)
	return c.cacheClient.Set(c.cartQuery.ctx, key, encoded, time.Hour).Err()
}

func (c CachedCartQuery) GetByUserID(userid uint32) (*[]Cart, error) {
	carts, err := c.GetFromCacheByUserID(userid)
	if err == nil {
		return carts, nil
	}

	carts, err = c.cartQuery.GetByUserID(userid)
	if err != nil {
		return nil, err
	}

	if err := c.SetCacheByUserID(userid, carts); err != nil {
		return nil, err
	}

	return carts, nil
}
