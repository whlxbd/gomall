package service

import (
	"context"
	"strconv"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/metadata"
	"github.com/whlxbd/gomall/app/checkout/infra/rpc"

	"github.com/whlxbd/gomall/common/utils/authpayload"
	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/cart"
	checkout "github.com/whlxbd/gomall/rpc_gen/kitex_gen/checkout"
	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/order"
	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/payment"
	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/product"

	_ "github.com/cloudwego/kitex/pkg/remote/codec/protobuf/encoding/gzip"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// Finish your business logic.
	// 鉴权
	token, err := authpayload.Token(s.ctx)
	if err != nil {
		klog.Errorf("get token failed: %v", err)
		return nil, kerrors.NewBizStatusError(401, "get token failed")
	}

	s.ctx = metadata.AppendToOutgoingContext(s.ctx, "Authorization", "Bearer "+token)

	payload, err := authpayload.Get(s.ctx)
	if err != nil {
		klog.Errorf("get payload failed: %v", err)
		return nil, kerrors.NewBizStatusError(401, "get payload failed")
	}

	klog.Infof("payload: %v", payload)

	if payload.UserId != int32(req.UserId) {
		return nil, kerrors.NewBizStatusError(401, "permission denied")
	}

	// 获取购物车
	cartResult, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: uint32(payload.UserId)})
	if err != nil {
		klog.Errorf("get cart failed: %v", err)
		return nil, kerrors.NewBizStatusError(400, "get cart failed")
	}

	if len(cartResult.Cart.Items) == 0 {
		klog.Errorf("cart is empty")
		return nil, kerrors.NewBizStatusError(400, "cart is empty")
	}

	itemIds := make([]uint32, 0)
	for _, item := range cartResult.Cart.Items {
		itemIds = append(itemIds, item.ProductId)
	}

	// 获取商品信息
	productsResult, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Ids: itemIds})
	if err != nil {
		klog.Errorf("get products failed: %v", err)
		return nil, kerrors.NewBizStatusError(400, "get products failed")
	}

	// 计算总价
	amount := float32(0)
	ois := []*order.OrderItem{}
	for id, pdt := range productsResult.Products {
		cost := pdt.Price * float32(cartResult.Cart.Items[id].Quantity)
		amount += cost
		ois = append(ois, &order.OrderItem{
			Item: &cart.CartItem{ProductId: pdt.Id, Quantity: cartResult.Cart.Items[id].Quantity},
			Cost: cost,
		})
	}

	// 下单
	placeOrderReq := order.PlaceOrderReq{
		UserId:       uint32(payload.UserId),
		UserCurrency: "CNY",
		Email:        req.Email,
		OrderItems:   ois,
	}

	if req.Address != nil {
		ZipCodeInt, err := strconv.Atoi(req.Address.ZipCode)
		if err != nil {
			klog.Errorf("invalid zip code: %v", err)
			return nil, kerrors.NewBizStatusError(400, "invalid zip code")
		}
		placeOrderReq.Address = &order.Address{
			StreetAddress: req.Address.StreetAddress,
			City:          req.Address.City,
			State:         req.Address.State,
			Country:       req.Address.Country,
			ZipCode:       int32(ZipCodeInt),
		}
	}

	placeOrderResult, err := rpc.OrderClient.PlaceOrder(s.ctx, &placeOrderReq)
	if err != nil {
		klog.Errorf("place order failed: %v", err)
		return nil, kerrors.NewBizStatusError(400, "place order failed")
	}

	orderId := placeOrderResult.Order.OrderId
	paymentReq := payment.ChargeReq{
		Amount:     amount,
		CreditCard: req.CreditCard,
		OrderId:    orderId,
		UserId:     uint32(payload.UserId),
	}

	// 支付
	paymentResult, err := rpc.PaymentClient.Charge(s.ctx, &paymentReq)
	if err != nil {
		klog.Errorf("payment failed: %v", err)
		return nil, kerrors.NewBizStatusError(400, "payment failed")
	}

	resp = &checkout.CheckoutResp{
		OrderId:       orderId,
		TransactionId: paymentResult.TransactionId,
	}

	return resp, nil
}
