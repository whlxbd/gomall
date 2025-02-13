package service

import (
	"context"
	"strconv"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/whlxbd/gomall/app/checkout/infra/rpc"

	"github.com/whlxbd/gomall/common/utils/authpayload"
	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/cart"
	checkout "github.com/whlxbd/gomall/rpc_gen/kitex_gen/checkout"
	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/order"
	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/payment"
	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/product"
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
	payload, err := authpayload.Get(s.ctx)
	if err != nil {
		klog.Errorf("get payload failed: %v", err)
		return nil, err
	}

	klog.Infof("payload: %v", payload)

	if payload.UserId == 0 {
		return nil, kerrors.NewBizStatusError(401, "user not found")
	}

	cartResult, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: uint32(payload.UserId)})
	if err != nil {
		return nil, err
	}

	if len(cartResult.Cart.Items) == 0 {
		return nil, kerrors.NewBizStatusError(400, "cart is empty")
	}

	itemIds := make([]uint32, 0)
	for _, item := range cartResult.Cart.Items {
		itemIds = append(itemIds, item.ProductId)
	}

	productsResult, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Ids: itemIds})
	if err != nil {
		return nil, err
	}

	amount := float32(0)
	ois := []*order.OrderItem{}
	for _, pdt := range productsResult.Products {
		cost := pdt.Price * float32(cartResult.Cart.Items[pdt.Id].Quantity)
		amount += cost
		ois = append(ois, &order.OrderItem{
			Item: &cart.CartItem{ProductId: pdt.Id, Quantity: cartResult.Cart.Items[pdt.Id].Quantity},
			Cost: cost,
		})
	}

	placeOrderReq := order.PlaceOrderReq{
		UserId:       uint32(payload.UserId),
		UserCurrency: "CNY",
		Email:        req.Email,
		OrderItems:   ois,
	}

	if req.Address != nil {
		ZipCodeInt, err := strconv.Atoi(req.Address.ZipCode)
		if err != nil {
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
		return nil, err
	}

	orderId := placeOrderResult.Order.OrderId
	paymentReq := payment.ChargeReq{
		Amount:     amount,
		CreditCard: req.CreditCard,
		OrderId:    orderId,
		UserId:     uint32(payload.UserId),
	}

	paymentResult, err := rpc.PaymentClient.Charge(s.ctx, &paymentReq)
	if err != nil {
		return nil, err
	}

	resp = &checkout.CheckoutResp{
		OrderId: orderId,
		TransactionId: paymentResult.TransactionId,
	}

	return resp, nil
}
