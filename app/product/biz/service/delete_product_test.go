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



func TestDeleteProduct_Run(t *testing.T) {
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
        req     *product.DeleteProductReq
        wantErr bool
        errCode int32
    }{
        {
            name: "管理员删除商品成功",
            ctx:  context.WithValue(context.Background(), "token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4iLCJhZG1pbiI6dHJ1ZX0.XbZL4omyujc11fq0Bki-l2cDoa7LKNRme3AmUdlr4_w"),
            req: &product.DeleteProductReq{
                Id: uint32(testProduct.ID),
            },
            wantErr: false,
        },
        // {
        //     name: "非管理员删除失败",
        //     ctx:  context.WithValue(context.Background(), "token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoidXNlciIsImFkbWluIjp0cnVlfQ.aP83FyfTZAEAkz-Nr1yE_Dsnz-CgDTQWA0bWb4Z6qIA"),
        //     req: &product.DeleteProductReq{
        //         Id: uint32(testProduct.ID),
        //     },
        //     wantErr: true,
        //     errCode: 400,
        // },
        {
            name: "商品ID不存在",
            ctx:  context.WithValue(context.Background(), "token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4iLCJhZG1pbiI6dHJ1ZX0.XbZL4omyujc11fq0Bki-l2cDoa7LKNRme3AmUdlr4_w"),
            req: &product.DeleteProductReq{
                Id: 9999,
            },
            wantErr: true,
            errCode: 400,
        },
        {
            name: "商品ID为0",
            ctx:  context.WithValue(context.Background(), "token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4iLCJhZG1pbiI6dHJ1ZX0.XbZL4omyujc11fq0Bki-l2cDoa7LKNRme3AmUdlr4_w"),
            req: &product.DeleteProductReq{
                Id: 0,
            },
            wantErr: true,
            errCode: 400,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            s := NewDeleteProductService(tt.ctx)
            resp, err := s.Run(tt.req)
            if (err != nil) != tt.wantErr {
                t.Errorf("DeleteProductService.Run() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if tt.wantErr {
                if bizErr, ok := kerrors.FromBizStatusError(err); ok && bizErr.BizStatusCode() != tt.errCode {
                    t.Errorf("DeleteProductService.Run() error code = %v, wantErr %v", bizErr.BizStatusCode(), tt.errCode)
                    return
                }
            }
            t.Logf("resp: %v", resp)
        })
    }
}