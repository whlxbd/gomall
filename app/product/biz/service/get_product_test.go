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

func TestGetProduct_Run(t *testing.T) {
    setupTestDB(t)
    
    // 预先创建测试商品
    testProducts := []*model.Product{
        {
            Name:        "测试商品",
            Description: "测试描述",
            Price:       99.9,
            Stock:       100,
            Status:      model.ProductStatusOnSale,
            Categories:  []model.Category{{Name: "测试分类"}},
        },
        {
            Name:        "测试商品2",
            Description: "测试描述2",
            Price:       199.9,
            Stock:       200,
            Status:      model.ProductStatusOnSale,
            Categories:  []model.Category{{Name: "测试分类2"}, {Name: "测试分类"}},
        },
    }
    for _, p := range testProducts {
        err := model.CreateProduct(mysql.DB, redis.RedisClient, context.Background(), p)
        assert.NoError(t, err)
    }

    tests := []struct {
        name    string
        req     *product.GetProductReq
        wantErr bool
        errCode int32
    }{
        {
            name: "获取商品成功",
            req: &product.GetProductReq{
                Ids: []uint32{1, 2},
            },
            wantErr: false,
        },
        {
            name: "商品不存在",
            req: &product.GetProductReq{
                Ids: []uint32{9999},
            },
            wantErr: true,
            errCode: 400,
        },
        {
            name: "商品ID为空",
            req: &product.GetProductReq{
                Ids: []uint32{},
            },
            wantErr: true,
            errCode: 400,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            s := NewGetProductService(context.Background())
            resp, err := s.Run(tt.req)
            if (err != nil) != tt.wantErr {
                t.Errorf("GetProductService.Run() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if tt.wantErr {
                if bizErr, ok := kerrors.FromBizStatusError(err); ok && bizErr.BizStatusCode() != tt.errCode {
                    t.Errorf("GetProductService.Run() error code = %v, wantErr %v", bizErr.BizStatusCode(), tt.errCode)
                    return
                }
            }
            if !tt.wantErr {
                for i, p := range resp.Products {
                    assert.Equal(t, testProducts[i].Name, p.Name)
                    assert.Equal(t, testProducts[i].Description, p.Description)
                    assert.Equal(t, testProducts[i].Price, p.Price)
                    assert.Equal(t, testProducts[i].Stock, p.Stock)
                    assert.Equal(t, len(testProducts[i].Categories), len(p.Categories))
                }
            }
            t.Logf("resp: %v", resp)
        })
    }
}