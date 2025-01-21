package service

import (
	"context"
	product "github.com/whlxbd/gomall/rpc_gen/kitex_gen/product"
	"github.com/whlxbd/gomall/app/product/biz/dal/model"
	"github.com/whlxbd/gomall/app/product/biz/dal/mysql"
	"github.com/whlxbd/gomall/app/product/biz/dal/redis"

	"github.com/cloudwego/kitex/pkg/kerrors"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	// Finish your business logic.
	if req.Id == 0 {
		return nil, kerrors.NewBizStatusError(400, "id is required")
	}

	p, err := model.NewCachedProductQuery(model.NewProductQuery(s.ctx, mysql.DB), redis.RedisClient).GetById(req.Id)
	if err != nil {
		return nil, err
	}

	return &product.GetProductResp{
		Product: &product.Product{
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
		},
	}, err

	// resp = &product.GetProductResp{
	// 	Product: &product.Product{
	// 		Id:          p.ID,
	// 		Name:        p.Name,
	// 		Description: p.Description,
	// 		Picture:     p.Picture,
	// 		Price:       p.Price,
	// 		Stock:       p.Stock,
	// 		Soldcount:   p.SoldCount,
	// 		Status:      product.ProductStatus(p.Status),
	// 		Ishot:       p.IsHot,
	// 		Isnew:       p.IsNew,
	// 		Isrecommend: p.IsRecommend,
	// 	},
	// }
	// return
}
