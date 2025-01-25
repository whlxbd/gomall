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

func TestListProducts_Run(t *testing.T) {
    setupTestDB(t)
    
    // 预先创建测试商品
    testProducts := []*model.Product{
        {
            Name:        "测试商品1",
            Description: "测试描述1",
            Price:       99.9,
            Stock:       100,
            Status:      model.ProductStatusOnSale,
            Categories:  []model.Category{{Name: "测试分类1"}},
        },
        {
            Name:        "测试商品2",
            Description: "测试描述2",
            Price:       199.9,
            Stock:       200,
            Status:      model.ProductStatusOnSale,
            Categories:  []model.Category{{Name: "测试分类2"}},
        },
    }
    
    for _, p := range testProducts {
        err := model.CreateProduct(mysql.DB, redis.RedisClient, context.Background(), p)
        assert.NoError(t, err)
    }

    tests := []struct {
        name        string
        req         *product.ListProductsReq
        wantErr     bool
        errCode     int32
        productLen  int
    }{
        {
            name: "查询所有商品",
            req: &product.ListProductsReq{
                Page:     1,
                PageSize: 100,
            },
            wantErr:    false,
            productLen: 2,
        },
        {
            name: "按分类查询商品",
            req: &product.ListProductsReq{
                CategoryName: "测试分类1",
                Page:        1,
                PageSize:    10,
            },
            wantErr:    false,
            productLen: 1,
        },
        {
            name: "查询不存在的分类",
            req: &product.ListProductsReq{
                CategoryName: "不存在的分类",
                Page:        1,
                PageSize:    10,
            },
            wantErr: true,
            errCode: 400,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            s := NewListProductsService(context.Background())
            resp, err := s.Run(tt.req)
            if (err != nil) != tt.wantErr {
                t.Errorf("ListProductsService.Run() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if tt.wantErr {
                if bizErr, ok := kerrors.FromBizStatusError(err); ok && bizErr.BizStatusCode() != tt.errCode {
                    t.Errorf("ListProductsService.Run() error code = %v, wantErr %v", bizErr.BizStatusCode(), tt.errCode)
                }
                return
            }
            assert.Equal(t, tt.productLen, len(resp.Products))
            t.Logf("resp: %+v", resp)
        })
    }
}