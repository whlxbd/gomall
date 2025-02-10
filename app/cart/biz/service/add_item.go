package service

import (
	"context"
	"fmt"

	"github.com/whlxbd/gomall/app/cart/biz/dal/model"
	"github.com/whlxbd/gomall/app/cart/biz/dal/mysql"
	"github.com/whlxbd/gomall/app/cart/biz/dal/redis"
	"github.com/whlxbd/gomall/app/cart/infra/rpc"
	cart "github.com/whlxbd/gomall/rpc_gen/kitex_gen/cart"
	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/whlxbd/gomall/common/utils/pool"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	// Finish your business logic.
	_ = pool.Submit(func() {
		klog.Infof("add item: %+v", req)
	})

	getProduct, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Ids: []uint32{req.Item.ProductId}})
	if err != nil {
		fmt.Printf("get product failed: %+v", err)
		_ = pool.Submit(func() {
			klog.Errorf("get product failed: %+v", err)
		})
		return resp, kerrors.NewBizStatusError(40000, "获取商品信息失败")
	}

	if len(getProduct.Products) == 0 || getProduct.Products == nil {
		_ = pool.Submit(func() {
			klog.Errorf("product not found: %+v", req.Item.ProductId)
		})
		return resp, kerrors.NewBizStatusError(40004, "商品不存在")
	}

	err = model.AddCart(&model.Cart{
		UserID:    req.UserId,
		ProductID: req.Item.ProductId,
		Quantity:  req.Item.Quantity,
		Selected:  true,
		Status:    true,
	}, mysql.DB, s.ctx, redis.RedisClient)
	if err != nil {
		_ = pool.Submit(func() {
			klog.Errorf("add cart failed: %+v", err)
		})
		return resp, kerrors.NewBizStatusError(40008, "添加购物车失败")
	}

	return
}
