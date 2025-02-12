package service

import (
	"context"

	"github.com/whlxbd/gomall/common/utils/pool"
	product "github.com/whlxbd/gomall/rpc_gen/kitex_gen/product"
	"github.com/whlxbd/gomall/app/product/biz/dal/mysql"
	"github.com/whlxbd/gomall/app/product/biz/dal/redis"
	"github.com/whlxbd/gomall/app/product/biz/dal/model"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type UpdateBatchProductService struct {
	ctx context.Context
} // NewUpdateBatchProductService new UpdateBatchProductService
func NewUpdateBatchProductService(ctx context.Context) *UpdateBatchProductService {
	return &UpdateBatchProductService{ctx: ctx}
}

// Run create note info
func (s *UpdateBatchProductService) Run(req *product.UpdateBatchProductReq) (resp *product.UpdateBatchProductResp, err error) {
	// Finish your business logic.
	_ = pool.Submit(func() {
		klog.Infof("UpdateBatchProductService Run %+v", req)
	})

	resp = &product.UpdateBatchProductResp{}

	if len(req.Products) == 0 {
		return nil, kerrors.NewBizStatusError(400, "products is required")
	}

	items := make([]*model.ProductStockItem, 0, len(req.Products))

	for _, p := range req.Products {
		if p.ProductId == 0 {
			_ = pool.Submit(func() {
				klog.Errorf("id is required: %+v", p)
			})
		}

		items = append(items, &model.ProductStockItem{
			ProductId: p.ProductId,
			StockChange: p.StockChange,
			SoldcountChange: p.SoldcountChange,
		})
	}

	err = model.UpdateBatchProduct(mysql.DB, redis.RedisClient, &s.ctx, items, req.IsStock)

	return
}
