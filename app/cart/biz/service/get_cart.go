package service

import (
	"context"

	cart "github.com/whlxbd/gomall/rpc_gen/kitex_gen/cart"
	"github.com/whlxbd/gomall/app/cart/biz/dal/mysql"
	"github.com/whlxbd/gomall/app/cart/biz/dal/redis"
	"github.com/whlxbd/gomall/app/cart/biz/dal/model"
	"github.com/whlxbd/gomall/common/utils/pool"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	// Finish your business logic.

	_ = pool.Submit(func() {
		klog.Infof("get cart: %+v", req)
	})


	carts, err := model.GetCartByUserId(req.UserId, mysql.DB, s.ctx, redis.RedisClient)
	if err != nil {
		_ = pool.Submit(func() {
			klog.Errorf("get cart failed: %+v", err)
		})
		return resp, kerrors.NewBizStatusError(40006, "获取购物车失败")
	}

	if len(carts) == 0  || carts == nil {
		_ = pool.Submit(func() {
			klog.Errorf("cart not found: %+v", req.UserId)
		})
		return resp, kerrors.NewBizStatusError(40007, "购物车为空")
	}

	var items []*cart.CartItem
	for _, c := range carts {
		items = append(items, &cart.CartItem{
			ProductId: c.ProductID,
			Quantity:  c.Quantity,
		})
	}

	return &cart.GetCartResp{
		Cart: &cart.Cart{
			UserId: req.UserId,
			Items: items,
		},
	}, nil
}
