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

// Create implements the RuleServiceImpl interface.
func (s *RuleServiceImpl) Create(ctx context.Context, req *auth.CreateReq) (resp *auth.CreateResp, err error) {
	resp, err = service.NewCreateService(ctx).Run(req)

	return resp, err
}

// List implements the RuleServiceImpl interface.
func (s *RuleServiceImpl) List(ctx context.Context, req *auth.ListReq) (resp *auth.ListResp, err error) {
	resp, err = service.NewListService(ctx).Run(req)

	return resp, err
}

// Delete implements the RuleServiceImpl interface.
func (s *RuleServiceImpl) Delete(ctx context.Context, req *auth.DeleteReq) (resp *auth.DeleteResp, err error) {
	resp, err = service.NewDeleteService(ctx).Run(req)

	return resp, err
}

// Get implements the RuleServiceImpl interface.
func (s *RuleServiceImpl) Get(ctx context.Context, req *auth.GetReq) (resp *auth.GetResp, err error) {
	resp, err = service.NewGetService(ctx).Run(req)

	return resp, err
}

// Update implements the RuleServiceImpl interface.
func (s *RuleServiceImpl) Update(ctx context.Context, req *auth.UpdateReq) (resp *auth.UpdateResp, err error) {
	resp, err = service.NewUpdateService(ctx).Run(req)

	return resp, err
}
