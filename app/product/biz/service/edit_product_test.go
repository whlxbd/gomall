package service

import (
    "context"
    "testing"
    "github.com/stretchr/testify/assert"
    product "github.com/whlxbd/gomall/rpc_gen/kitex_gen/product"
    "github.com/whlxbd/gomall/app/product/biz/dal/mysql"
    "github.com/whlxbd/gomall/app/product/biz/dal/redis"
    "github.com/whlxbd/gomall/app/product/biz/dal/model"
    "github.com/cloudwego/kitex/pkg/kerrors"
)

func TestEditProduct_Run(t *testing.T) {
    setupTestDB(t)
    
    // 预先创建一个测试商品
    testProduct := &model.Product{
        Name:        "测试商品",
        Description: "测试描述",
        Price:       99.9,
        Stock:       100,
        Status:      model.ProductStatusOnSale,
        Categories:  []model.Category{{Name: "测试分类"}},
    }
    err := model.CreateProduct(mysql.DB, redis.RedisClient, context.Background(), testProduct)
    assert.NoError(t, err)

    tests := []struct {
        name    string
        ctx     context.Context
        req     *product.EditProductReq
        wantErr bool
        errCode int32
    }{
        {
            name: "管理员修改商品成功",
            ctx:  context.WithValue(context.Background(), "token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4iLCJhZG1pbiI6dHJ1ZX0.XbZL4omyujc11fq0Bki-l2cDoa7LKNRme3AmUdlr4_w"),
            req: &product.EditProductReq{
                Id:          uint32(testProduct.ID),
                Name:        "修改后的商品",
                Description: "修改后的描述",
                Price:       199.9,
                Stock:       50,
                Categories:  []string{"修改后的分类"},
                Status:      2,
            },
            wantErr: false,
        },
        {
            name: "商品不存在",
            ctx:  context.WithValue(context.Background(), "token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4iLCJhZG1pbiI6dHJ1ZX0.XbZL4omyujc11fq0Bki-l2cDoa7LKNRme3AmUdlr4_w"),
            req: &product.EditProductReq{
                Id:          9999,
                Name:        "测试商品",
                Description: "测试描述",
                Price:       99.9,
                Stock:       100,
                Categories:  []string{"测试分类"},
                Status:      1,
            },
            wantErr: true,
            errCode: 400,
        },
        {
            name: "商品参数验证失败-名称为空",
            ctx:  context.WithValue(context.Background(), "token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4iLCJhZG1pbiI6dHJ1ZX0.XbZL4omyujc11fq0Bki-l2cDoa7LKNRme3AmUdlr4_w"),
            req: &product.EditProductReq{
                Id:          uint32(testProduct.ID),
                Name:        "",
                Description: "测试描述",
                Price:       99.9,
                Stock:       100,
                Categories:  []string{"测试分类"},
                Status:      1,
            },
            wantErr: true,
            errCode: 400,
        },
        {
            name: "商品参数验证失败-价格小于0",
            ctx:  context.WithValue(context.Background(), "token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4iLCJhZG1pbiI6dHJ1ZX0.XbZL4omyujc11fq0Bki-l2cDoa7LKNRme3AmUdlr4_w"),
            req: &product.EditProductReq{
                Id:          uint32(testProduct.ID),
                Name:        "测试商品",
                Description: "测试描述",
                Price:       -1,
                Stock:       100,
                Categories:  []string{"测试分类"},
                Status:      1,
            },
            wantErr: true,
            errCode: 400,
        },
        {
            name: "商品ID为0",
            ctx:  context.WithValue(context.Background(), "token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4iLCJhZG1pbiI6dHJ1ZX0.XbZL4omyujc11fq0Bki-l2cDoa7LKNRme3AmUdlr4_w"),
            req: &product.EditProductReq{
                Id:          0,
                Name:        "测试商品",
                Description: "测试描述",
                Price:       99.9,
                Stock:       100,
                Categories:  []string{"测试分类"},
                Status:      1,
            },
            wantErr: true,
            errCode: 400,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            s := NewEditProductService(tt.ctx)
            resp, err := s.Run(tt.req)
            if (err != nil) != tt.wantErr {
                t.Errorf("EditProductService.Run() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if tt.wantErr {
                if bizErr, ok := kerrors.FromBizStatusError(err); ok && bizErr.BizStatusCode() != tt.errCode {
                    t.Errorf("EditProductService.Run() error code = %v, wantErr %v", bizErr.BizStatusCode(), tt.errCode)
                    return
                }
            }
            t.Logf("resp: %v", resp)
        })
    }
}