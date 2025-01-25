package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/whlxbd/gomall/app/product/biz/dal/model"
	product "github.com/whlxbd/gomall/rpc_gen/kitex_gen/product"

	// "github.com/whlxbd/gomall/app/product/biz/middleware"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/whlxbd/gomall/app/product/biz/dal/mysql"
	"github.com/whlxbd/gomall/app/product/biz/dal/redis"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// Finish your business logic.
	// if middleware.CheckAdminPermission(s.ctx) != nil {
	// 	return nil, kerrors.NewBizStatusError(400, "permission denied: admin required")
	// }

	klog.Infof("ListProductsService input: %+v\n", req)

	resp = &product.ListProductsResp{}
	products := new([]model.Product)

	if req.Page <= 0 {
		req.Page = 1
	}

	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	if req.CategoryName == "" {
		products, err = model.GetByCategoryName(mysql.DB, &s.ctx, 0, req.Page, req.PageSize)
		if err != nil {
			klog.Error("查询所有商品列表失败", err)
			return nil, kerrors.NewBizStatusError(400, "查询所有商品列表失败")
		}
	} else {
		ccq := model.NewCachedCategoryQuery(model.NewCategoryQuery(s.ctx, mysql.DB), redis.RedisClient)

		id, err := ccq.IsExistByName(req.CategoryName)
		if err != nil {
			klog.Error("查询分类失败", err)
			return nil, kerrors.NewBizStatusError(400, "查询分类失败")
		}

		if id == 0 {
			klog.Error("分类不存在", err)
			return nil, kerrors.NewBizStatusError(400, "分类不存在")
		}

		products, err = model.GetByCategoryName(mysql.DB, &s.ctx, id, req.Page, req.PageSize)
		if err != nil {
			klog.Error("查询商品列表失败", err)
			return nil, kerrors.NewBizStatusError(400, "查询商品列表失败")
		}
	}

	for _, p := range *products {
		categories := make([]string, len(p.Categories))
		for i, cat := range p.Categories {
			categories[i] = cat.Name
		}
		resp.Products = append(resp.Products, &product.Product{
			Id:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Picture:     p.Picture,
			Stock:       p.Stock,
			Soldcount:   p.SoldCount,
			Status:      product.ProductStatus(p.Status),
			Ishot: 	 	 p.IsHot,
			Isnew: 	 	 p.IsNew,
			Isrecommend: p.IsRecommend,
			Categories:  categories,
		})
	}

	return
}
