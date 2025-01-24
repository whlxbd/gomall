package service

import (
	"context"
	product "github.com/whlxbd/gomall/rpc_gen/kitex_gen/product"
	"github.com/whlxbd/gomall/app/product/biz/dal/model"
	// "github.com/whlxbd/gomall/app/product/biz/middleware"
	"github.com/whlxbd/gomall/app/product/biz/dal/mysql"
	"github.com/whlxbd/gomall/app/product/biz/dal/redis"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
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
	// if middleware.CheckAdminPermission(s.ctx) != nil {
	// 	return nil, kerrors.NewBizStatusError(400, "permission denied: admin required")
	// }

	klog.Infof("DeleteProductService input: %+v\n", req)

	resp = &product.DeleteProductResp{}

	if req.Id == 0 {
		resp.Success = false
		return resp, kerrors.NewBizStatusError(400, "id is required")
	}

	p, err := model.NewCachedProductQuery(model.NewProductQuery(s.ctx, mysql.DB), redis.RedisClient).GetById(req.Id)

	if err != nil { // 查询商品失败
		resp.Success = false
		return resp, err
	}

	if p == nil { // 商品不存在
		resp.Success = false
		klog.Errorf("product not found, id: %d\n", req.Id)
		return resp, kerrors.NewBizStatusError(400, "product not found")
	}

	err = model.DeleteProduct(mysql.DB, redis.RedisClient, s.ctx, req.Id)	// 删除商品
	if err != nil {
		klog.Errorf("delete product failed, id: %d, err: %v\n", req.Id, err)
		resp.Success = false
		return resp, kerrors.NewBizStatusError(400, "delete product failed")
	}

	resp.Success = true
	return
}
