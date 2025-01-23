package service

import (
	"context"
	product "github.com/whlxbd/gomall/rpc_gen/kitex_gen/product"
	"github.com/whlxbd/gomall/app/product/biz/dal/model"
	"github.com/whlxbd/gomall/app/product/biz/middleware"
	"github.com/whlxbd/gomall/app/product/biz/dal/mysql"
	"github.com/whlxbd/gomall/app/product/biz/dal/redis"

	"github.com/cloudwego/kitex/pkg/kerrors"
)

type CreateProductService struct {
	ctx context.Context
} // NewCreateProductService new CreateProductService
func NewCreateProductService(ctx context.Context) *CreateProductService {
	return &CreateProductService{ctx: ctx}
}

// Run create note info
func (s *CreateProductService) Run(req *product.CreateProductReq) (resp *product.CreateProductResp, err error) {
	// Finish your business logic.
	if err = middleware.CheckAdminPermission(s.ctx); err != nil {
		return nil, kerrors.NewBizStatusError(400, err.Error())
	}

	if req.Name == "" {
		return nil, kerrors.NewBizStatusError(400, "name is required")
	}

	if req.Price < 0 || req.Stock < 0  || req.Soldcount < 0 {
		return nil, kerrors.NewBizStatusError(400, "price, stock, soldcount, must be greater than 0")
	}

	if err := model.ProductStatus(req.Status).IsValid(); err != nil {
		return nil, kerrors.NewBizStatusError(400, "status is invalid")
	}

	categories := make([]model.Category, len(req.Categories))
	for i, cat := range req.Categories {
		categories[i] = model.Category{Name: cat}
	}
	p := &model.Product{
		Name:        req.Name,
		Description: req.Description,
		Picture:     req.Picture,
		Price:       req.Price,
		Stock:       req.Stock,
		SoldCount:   req.Soldcount,
		Status:      model.ProductStatus(req.Status),
		IsHot:       req.Ishot,
		IsNew:       req.Isnew,
		IsRecommend: req.Isrecommend,
		Categories:  categories,
	}

	if err := model.CreateProduct(mysql.DB, redis.RedisClient, s.ctx, p); err != nil {
		return nil, err
	}
	return &product.CreateProductResp{
		ProductId: p.ID,
		Success:  true,
	}, err
}
