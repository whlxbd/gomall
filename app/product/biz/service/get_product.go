package service

import (
	"context"

	"github.com/whlxbd/gomall/app/product/biz/dal/model"
	"github.com/whlxbd/gomall/app/product/biz/dal/mysql"
	"github.com/whlxbd/gomall/app/product/biz/dal/redis"
	product "github.com/whlxbd/gomall/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	// klog.Error("GetProduct Start")

	if len(req.Ids) == 0 {
		klog.Errorf("ids is required")
		return nil, kerrors.NewBizStatusError(400, "ids is required")
	}

	products, err := model.NewCachedProductQuery(
		model.NewProductQuery(s.ctx, mysql.DB),
		redis.RedisClient,
	).GetByIds(req.Ids)

	if err != nil {
		klog.Errorf("get product by ids failed: %v", err)
		return nil, err
	}

	if len(products) == 0 {
		klog.Errorf("products not found, ids: %v", req.Ids)
		return nil, kerrors.NewBizStatusError(400, "products not found")
	}

	respProducts := make([]*product.Product, 0, len(products))
	for _, p := range products {
		categories := make([]string, len(p.Categories))
		for i, cat := range p.Categories {
			categories[i] = cat.Name
		}

		respProducts = append(respProducts, &product.Product{
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
			Categories:  categories,
		})
	}

	return &product.GetProductResp{
		Products: respProducts,
	}, nil
}
