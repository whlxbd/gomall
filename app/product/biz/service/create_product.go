package service

import (
	"context"
	product "github.com/whlxbd/gomall/rpc_gen/kitex_gen/product"
	// "github.com/whlxbd/gomall/app/product/biz/dal/model"
	"github.com/whlxbd/gomall/app/product/biz/middleware"
	// "github.com/whlxbd/gomall/app/product/biz/dal/mysql"
	// "github.com/whlxbd/gomall/app/product/biz/dal/redis"

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
	if middleware.CheckAdminPermission(s.ctx) != nil {
		return nil, kerrors.NewBizStatusError(400, "permission denied: user role is not allowed")
	}

	
	return
}
