package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/whlxbd/gomall/common/utils/authpayload"
	checkout "github.com/whlxbd/gomall/rpc_gen/kitex_gen/checkout"
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

	return
}
