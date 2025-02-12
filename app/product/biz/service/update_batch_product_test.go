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

func TestUpdateBatchProduct_Run(t *testing.T) {
    setupTestDB(t)
    
    // 预先创建测试商品
    testProducts := []*model.Product{
        {
            Name:        "商品1",
            Description: "描述1",
            Price:       100,
            Stock:       50,
            Status:      model.ProductStatusOnSale,
        },
        {
            Name:        "商品2",
            Description: "描述2",
            Price:       200,
            Stock:       30,
            Status:      model.ProductStatusOnSale,
        },
    }
    
    for _, p := range testProducts {
        err := model.CreateProduct(mysql.DB, redis.RedisClient, context.Background(), p)
        assert.NoError(t, err)
    }

    tests := []struct {
        name    string
        req     *product.UpdateBatchProductReq
        wantErr bool
        errCode int32
    }{
        {
            name: "批量扣减库存成功",
            req: &product.UpdateBatchProductReq{
                Products: []*product.ProductBatch{
                    {
                        ProductId:    uint32(testProducts[0].ID),
                        StockChange:  -10,
                    },
                    {
                        ProductId:    uint32(testProducts[1].ID),
                        StockChange:  -5,
                    },
                },
                IsStock: true,
            },
            wantErr: false,
        },
        {
            name: "批量更新销量成功",
            req: &product.UpdateBatchProductReq{
                Products: []*product.ProductBatch{
                    {
                        ProductId:       uint32(testProducts[0].ID),
                        SoldcountChange: 10,
                    },
                    {
                        ProductId:       uint32(testProducts[1].ID),
                        SoldcountChange: 5,
                    },
                },
                IsStock: false,
            },
            wantErr: false,
        },
        {
            name: "库存不足失败",
            req: &product.UpdateBatchProductReq{
                Products: []*product.ProductBatch{
                    {
                        ProductId:    uint32(testProducts[0].ID),
                        StockChange:  -100,
                    },
                },
                IsStock: true,
            },
            wantErr: true,
            errCode: 40010,
        },
        {
            name: "商品列表为空",
            req: &product.UpdateBatchProductReq{
                Products: []*product.ProductBatch{},
                IsStock: true,
            },
            wantErr: true,
            errCode: 400,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            s := NewUpdateBatchProductService(context.Background())
            resp, err := s.Run(tt.req)
            
            if (err != nil) != tt.wantErr {
                t.Errorf("UpdateBatchProductService.Run() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            
            if tt.wantErr {
                if bizErr, ok := kerrors.FromBizStatusError(err); ok && bizErr.BizStatusCode() != tt.errCode {
                    t.Errorf("UpdateBatchProductService.Run() error code = %v, wantErr %v", bizErr.BizStatusCode(), tt.errCode)
                    return
                }
            } else {
                // 验证库存/销量是否正确更新
                for i, p := range tt.req.Products {
                    var updatedProduct model.Product
                    err := mysql.DB.First(&updatedProduct, p.ProductId).Error
                    assert.NoError(t, err)
                    
                    if tt.req.IsStock {
                        assert.Equal(t, testProducts[i].Stock+p.StockChange, updatedProduct.Stock)
                    } else {
                        assert.Equal(t, testProducts[i].SoldCount+p.SoldcountChange, updatedProduct.SoldCount)
                    }
                }
            }
            
            t.Logf("resp: %v", resp)
        })
    }
}