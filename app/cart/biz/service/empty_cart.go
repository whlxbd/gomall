package service

import (
	"context"

	cart "github.com/whlxbd/gomall/rpc_gen/kitex_gen/cart"
	"github.com/whlxbd/gomall/app/cart/biz/dal/mysql"
	"github.com/whlxbd/gomall/app/cart/biz/dal/redis"
	"github.com/whlxbd/gomall/app/cart/biz/dal/model"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/whlxbd/gomall/common/utils/pool"
)

type EmptyCartService struct {
	ctx context.Context
} // NewEmptyCartService new EmptyCartService
func NewEmptyCartService(ctx context.Context) *EmptyCartService {
	return &EmptyCartService{ctx: ctx}
} 

// Run create note info
func (s *EmptyCartService) Run(req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	// Finish your business logic.
	_ = pool.Submit(func() {
		klog.Infof("empty cart: %+v", req)
	})

	err = model.EmptyCart(req.UserId, mysql.DB, s.ctx, redis.RedisClient)
	if err != nil {
		_ = pool.Submit(func() {
			klog.Errorf("empty cart failed: %+v", err)
		})
		return resp, kerrors.NewBizStatusError(40005, "清空购物车失败")
	}
	return
}
