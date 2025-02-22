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

// CheckPermission implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) CheckPermission(ctx context.Context, req *auth.CheckPermissionReq) (resp *auth.CheckPermissionResp, err error) {
	resp, err = service.NewCheckPermissionService(ctx).Run(req)

	return resp, err
}

// CheckWhite implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) CheckWhite(ctx context.Context, req *auth.CheckWhiteReq) (resp *auth.CheckWhiteResp, err error) {
	resp, err = service.NewCheckWhiteService(ctx).Run(req)

	return resp, err
}

// LoadPolicy implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) LoadPolicy(ctx context.Context, req *auth.LoadPolicyReq) (resp *auth.LoadPolicyResp, err error) {
	resp, err = service.NewLoadPolicyService(ctx).Run(req)

	return resp, err
}

// RemovePolicy implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) RemovePolicy(ctx context.Context, req *auth.RemovePolicyReq) (resp *auth.RemovePolicyResp, err error) {
	resp, err = service.NewRemovePolicyService(ctx).Run(req)

	return resp, err
}
