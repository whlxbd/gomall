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

type DeleteProductService struct {
	ctx context.Context
} // NewDeleteProductService new DeleteProductService
func NewDeleteProductService(ctx context.Context) *DeleteProductService {
	return &DeleteProductService{ctx: ctx}
}

// Run create note info
func (s *DeleteProductService) Run(req *product.DeleteProductReq) (resp *product.DeleteProductResp, err error) {
	// Finish your business logic.
	if middleware.CheckAdminPermission(s.ctx) != nil {
		return nil, kerrors.NewBizStatusError(400, "permission denied: admin required")
	}

	if req.Id == 0 {
		return nil, kerrors.NewBizStatusError(400, "id is required")
	}

	err = model.DeleteProduct(mysql.DB, redis.RedisClient, s.ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return
}
