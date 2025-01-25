package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/whlxbd/gomall/app/product/biz/dal/model"
	product "github.com/whlxbd/gomall/rpc_gen/kitex_gen/product"

	// "github.com/whlxbd/gomall/app/product/biz/middleware"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/whlxbd/gomall/app/product/biz/dal/mysql"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	// Finish your business logic.
	klog.Infof("SearchProductsService input: %+v\n", req)

	resp = &product.SearchProductsResp{}

	if req.Query == "" {
        return nil, kerrors.NewBizStatusError(400, "query is required")
    }
    if req.Page <= 0 {
        return nil, kerrors.NewBizStatusError(400, "page must be greater than 0")
    }
    if req.PageSize <= 0 {
        return nil, kerrors.NewBizStatusError(400, "page_size must be greater than 0")
    }

	products, err := model.SearchProduct(mysql.DB, &s.ctx, req.Query, req.Page, req.PageSize)

	if err != nil {
		klog.Error("查询商品失败", err)
		return nil, kerrors.NewBizStatusError(400, "查询商品失败")
	}

	resp.Results = make([]*product.Product, len(products))
	for i, p := range products {
		Categories := make([]string, len(p.Categories))
		for i, cat := range p.Categories {
			Categories[i] = cat.Name
		}
		resp.Results[i] = &product.Product{
			Id:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Picture:     p.Picture,
			Price:       p.Price,
			Stock:       p.Stock,
			Soldcount:   p.SoldCount,
			Status:      product.ProductStatus(p.Status),
			Ishot:       p.IsHot,
			Isnew:       p.IsNew,
			Isrecommend: p.IsRecommend,
			Categories:  Categories,
		}
	}
	return
}
