package main

import (
	"context"

	"github.com/whlxbd/gomall/app/auth/biz/service"
	auth "github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth"
)

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct{}

// DeliverTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) DeliverTokenByRPC(ctx context.Context, req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	resp, err = service.NewDeliverTokenByRPCService(ctx).Run(req)

	return resp, err
}

// VerifyTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) VerifyTokenByRPC(ctx context.Context, req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	resp, err = service.NewVerifyTokenByRPCService(ctx).Run(req)

	return resp, err
}

// GetPayload implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) GetPayload(ctx context.Context, req *auth.GetPayloadReq) (resp *auth.GetPayloadResp, err error) {
	resp, err = service.NewGetPayloadService(ctx).Run(req)

	return resp, err
}

// Authenticate implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) Authenticate(ctx context.Context, req *auth.AuthenticateReq) (resp *auth.AuthenticateResp, err error) {
	resp, err = service.NewAuthenticateService(ctx).Run(req)

	return resp, err
}
