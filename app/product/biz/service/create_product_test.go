package service

import (
	"context"
    "testing"
    "os"
    "github.com/stretchr/testify/assert"
    product "github.com/whlxbd/gomall/rpc_gen/kitex_gen/product"
    "github.com/whlxbd/gomall/app/product/biz/dal/mysql"
    "github.com/whlxbd/gomall/app/product/biz/dal/redis"
    "github.com/whlxbd/gomall/app/product/biz/dal/model"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    goredis "github.com/redis/go-redis/v9"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

func setupTestDB(t *testing.T) {
    os.Setenv("JWT_SECREAT", "test")
    db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    assert.NoError(t, err)
    
    err = db.AutoMigrate(&model.Product{}, &model.Category{})
    assert.NoError(t, err)
    
    mysql.DB = db
    
    redis.RedisClient = goredis.NewClient(&goredis.Options{
        Addr: "localhost:6379",
    })
}

func TestCreateProduct_Run(t *testing.T) {
	setupTestDB(t)
	tests := []struct {
        name    string
        ctx     context.Context
        req     *product.CreateProductReq
        wantErr bool
        errCode int32
    }{
        {
            name: "管理员创建商品成功",
            ctx:  context.WithValue(context.Background(), "token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4iLCJhZG1pbiI6dHJ1ZX0.XbZL4omyujc11fq0Bki-l2cDoa7LKNRme3AmUdlr4_w"),
            req: &product.CreateProductReq{
                Name:        "测试商品",
                Description: "商品描述",
                Price:      99.9,
                Stock:      100,
                Categories: []string{"测试分类"},
                Status:     1,
            },
            wantErr: false,
        },
        {
            name: "非管理员创建失败",
            ctx:  context.WithValue(context.Background(), "token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoidXNlciIsImFkbWluIjp0cnVlfQ.aP83FyfTZAEAkz-Nr1yE_Dsnz-CgDTQWA0bWb4Z6qIA"),
            req: &product.CreateProductReq{
                Name:        "测试商品",
                Description: "商品描述",
                Price:      99.9,
                Stock:      100,
                Categories: []string{"测试分类"},
                Status:     1,
            },
            wantErr: true,
            errCode: 400,
        },
        {
            name: "商品参数验证失败",
            ctx:  context.WithValue(context.Background(), "token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4iLCJhZG1pbiI6dHJ1ZX0.XbZL4omyujc11fq0Bki-l2cDoa7LKNRme3AmUdlr4_w"),
            req: &product.CreateProductReq{
                Name:        "",
                Description: "商品描述",
                Price:      99.9,
                Stock:      100,
                Categories: []string{"测试分类"},
                Status:     1,
            },
            wantErr: true,
            errCode: 400,
        },
        {
            name: "商品参数验证失败",
            ctx:  context.WithValue(context.Background(), "token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4iLCJhZG1pbiI6dHJ1ZX0.XbZL4omyujc11fq0Bki-l2cDoa7LKNRme3AmUdlr4_w"),
            req: &product.CreateProductReq{
                Name:        "测试商品",
                Description: "商品描述",
                Price:      -1,
                Stock:      100,
                Categories: []string{"测试分类"},
                Status:     1,
            },
            wantErr: true,
            errCode: 400,
        },
        {
            name: "商品参数验证失败",
            ctx:  context.WithValue(context.Background(), "token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4iLCJhZG1pbiI6dHJ1ZX0.XbZL4omyujc11fq0Bki-l2cDoa7LKNRme3AmUdlr4_w"),
            req: &product.CreateProductReq{
                Name:        "测试商品",
                Description: "商品描述",
                Price:      99.9,
                Stock:      100,
                Categories: []string{"测试分类"},
                Status:     0,
            },
            wantErr: true,
            errCode: 400,
        },
    }
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewCreateProductService(tt.ctx)
			resp, err := s.Run(tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateProductService.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr{
				if bizErr, ok := kerrors.FromBizStatusError(err); ok && bizErr.BizStatusCode() != tt.errCode {
					t.Errorf("CreateProductService.Run() error code = %v, wantErr %v", bizErr.BizStatusCode(), tt.errCode)
					return
				}
			}
			t.Logf("resp: %v", resp)
		})
	}
	// init req and assert value

	// req := &product.CreateProductReq{}
	// resp, err := s.Run(req)
	// t.Logf("err: %v", err)
	// t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
