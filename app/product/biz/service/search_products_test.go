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

func TestSearchProducts_Run(t *testing.T) {
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
            Name:        "苹果手机",
            Description: "iPhone测试",
            Price:       199.9,
            Stock:       200,
            Status:      model.ProductStatusOnSale,
            Categories:  []model.Category{{Name: "手机分类"}},
        },
    }
    
    for _, p := range testProducts {
        err := model.CreateProduct(mysql.DB, redis.RedisClient, context.Background(), p)
        assert.NoError(t, err)
    }

    tests := []struct {
        name       string
        req        *product.SearchProductsReq
        wantErr    bool
        errCode    int32
        resultSize int
    }{
        {
            name: "搜索成功",
            req: &product.SearchProductsReq{
                Query:    "测试",
                Page:     1,
                PageSize: 10,
            },
            wantErr:    false,
            resultSize: 2,
        },
        {
            name: "搜索手机",
            req: &product.SearchProductsReq{
                Query:    "手机",
                Page:     1,
                PageSize: 10,
            },
            wantErr:    false,
            resultSize: 1,
        },
        {
            name: "空查询",
            req: &product.SearchProductsReq{
                Query:    "",
                Page:     1,
                PageSize: 10,
            },
            wantErr: true,
            errCode: 400,
        },
        {
            name: "无结果搜索",
            req: &product.SearchProductsReq{
                Query:    "不存在的商品",
                Page:     1,
                PageSize: 10,
            },
            wantErr:    false,
            resultSize: 0,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            s := NewSearchProductsService(context.Background())
            resp, err := s.Run(tt.req)
            if (err != nil) != tt.wantErr {
                t.Errorf("SearchProductsService.Run() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if tt.wantErr {
                if bizErr, ok := kerrors.FromBizStatusError(err); ok && bizErr.BizStatusCode() != tt.errCode {
                    t.Errorf("SearchProductsService.Run() error code = %v, wantErr %v", bizErr.BizStatusCode(), tt.errCode)
                }
                return
            }
            assert.Equal(t, tt.resultSize, len(resp.Results))
            t.Logf("resp: %v", resp)
        })
    }
}