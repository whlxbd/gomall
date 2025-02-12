package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/metadata"
	"github.com/whlxbd/gomall/app/checkout/infra/rpc"
	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth"
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
	md, ok := metadata.FromIncomingContext(s.ctx)
	if !ok {
		return nil, kerrors.NewBizStatusError(400, "metadata not found")
	}
	tokens := md.Get("token")
	if len(tokens) == 0 || tokens[0] == "" {
		return nil, kerrors.NewBizStatusError(400, "token not found")
	}
	token := tokens[0]

	authVerifyResp, err := rpc.AuthClient.VerifyTokenByRPC(s.ctx, &auth.VerifyTokenReq{Token: token})
	if err != nil {
		return nil, kerrors.NewBizStatusError(500, err.Error())
	}
	if authVerifyResp.Res == false {
		return nil, kerrors.NewBizStatusError(401, "token invalid")
	}

	payload, err := rpc.AuthClient.GetPayload(s.ctx, &auth.GetPayloadReq{Token: token})
	if err != nil {
		return nil, kerrors.NewBizStatusError(500, err.Error())
	}

	if payload.UserId == 0 {
		return nil, kerrors.NewBizStatusError(401, "user not found")
	}

	return
}
