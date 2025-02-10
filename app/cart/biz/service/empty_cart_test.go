package service

import (
    "context"
    "testing"
	"os"

    "github.com/stretchr/testify/assert"
    "github.com/whlxbd/gomall/app/cart/biz/dal/model"
    "github.com/whlxbd/gomall/app/cart/biz/dal/mysql"
    "github.com/whlxbd/gomall/app/cart/biz/dal/redis"
    cart "github.com/whlxbd/gomall/rpc_gen/kitex_gen/cart"
	"github.com/whlxbd/gomall/common/utils/pool"
	goredis "github.com/redis/go-redis/v9"
	"gorm.io/driver/sqlite"
	"github.com/alicebob/miniredis/v2"
    "gorm.io/gorm"
)

func setupTestDB(t *testing.T) {
    os.Setenv("JWT_SECREAT", "test")
    db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    assert.NoError(t, err)
    
    err = db.AutoMigrate(&model.Cart{})
    assert.NoError(t, err)
    
    mysql.DB = db
    
	mr, err := miniredis.Run()
    assert.NoError(t, err)
    
    redis.RedisClient = goredis.NewClient(&goredis.Options{
        Addr: mr.Addr(),
    })
    pool.Init()
}

func TestEmptyCart_Run(t *testing.T) {
    setupTestDB(t)

    // 准备测试数据
    testCarts := []*model.Cart{
        {
            UserID:    1,
            ProductID: 1,
            Quantity:  1,
            Selected:  true,
            Status:    true,
        },
        {
            UserID:    1,
            ProductID: 2,
            Quantity:  2,
            Selected:  true,
            Status:    true,
        },
    }

    // 写入测试数据
    for _, cart := range testCarts {
        err := mysql.DB.Create(cart).Error
        assert.NoError(t, err)
    }

    tests := []struct {
        name    string
        req     *cart.EmptyCartReq
        wantErr bool
    }{
        {
            name: "清空购物车成功",
            req: &cart.EmptyCartReq{
                UserId: 1,
            },
            wantErr: false,
        },
        {
            name: "用户不存在",
            req: &cart.EmptyCartReq{
                UserId: 999,
            },
            wantErr: false,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            s := NewEmptyCartService(context.Background())
            resp, err := s.Run(tt.req)

            if (err != nil) != tt.wantErr {
                t.Errorf("EmptyCartService.Run() error = %v, wantErr %v", err, tt.wantErr)
                return
            }

            // 验证购物车是否清空
            var carts []*model.Cart
            err = mysql.DB.Where("user_id = ?", tt.req.UserId).Find(&carts).Error
            assert.NoError(t, err)
            assert.Empty(t, carts)

            // 验证缓存是否清空
            carts, err = model.GetCartByUserId(tt.req.UserId, mysql.DB, context.Background(), redis.RedisClient)
            assert.NoError(t, err)
            assert.Empty(t, carts)

            t.Logf("resp: %v", resp)
        })
    }
}