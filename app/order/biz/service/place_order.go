package service

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"github.com/whlxbd/gomall/app/order/biz/dal/model"
	"github.com/whlxbd/gomall/app/order/biz/dal/mq"
	"github.com/whlxbd/gomall/app/order/biz/dal/mysql"
	"github.com/whlxbd/gomall/common/utils/pool"
	order "github.com/whlxbd/gomall/rpc_gen/kitex_gen/order"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	// Finish your business logic.
	if len(req.OrderItems) == 0 {
		_ = pool.Submit(func() {
			klog.Errorf("order items is empty")
		})
	}

	orderId, _ := uuid.NewUUID()
	o := &model.Order{
		OrderId:      orderId.String(),
		OrderState:   model.OrderStatePlaced,
		UserId:       req.UserId,
		UserCurrency: req.UserCurrency,
		Consignee: model.Consignee{
			Email: req.Email,
		},
	}

	if req.Address != nil {
		a := req.Address
		o.Consignee.Country = a.Country
		o.Consignee.State = a.State
		o.Consignee.City = a.City
		o.Consignee.StreetAddress = a.StreetAddress
	}

	if err = model.CreateOrder(mysql.DB, s.ctx, o); err != nil {
		_ = pool.Submit(func() {
			klog.Errorf("model.CreateOrder.err:%v", err)
		})
		return
	}

	var items []model.OrderItem
	for _, v := range req.OrderItems {
		items = append(items, model.OrderItem{
			OrderIdRefer: orderId.String(),
			ProductId:    v.Item.ProductId,
			Quantity:     v.Item.Quantity,
			Cost:         v.Cost,
		})
	}

	if err = model.CreateOrderItems(mysql.DB, s.ctx, &items); err != nil {
		_ = pool.Submit(func() {
			klog.Errorf("model.CreateOrderItems.err:%v", err)
		})
		return
	}

	_ = pool.Submit(func() {
		err2 := mq.SendOrderTimeoutMessage(context.Background(), orderId.String(), req.UserId)
		if err2 != nil {
			fmt.Printf("mq.SendOrderTimeoutMessage failed: %+v\n", err2)
			klog.Errorf("mq.SendOrderTimeoutMessage.err:%v", err2)
		} else {
			fmt.Printf("mq.SendOrderTimeoutMessage success\n")
		}
	})

	resp = &order.PlaceOrderResp{
		Order: &order.OrderResult{
			OrderId: orderId.String(),
		},
	}
	return
}
