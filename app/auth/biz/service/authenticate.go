package service

import (
	"context"
	auth "github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth"
)

type AuthenticateService struct {
	ctx context.Context
} // NewAuthenticateService new AuthenticateService
func NewAuthenticateService(ctx context.Context) *AuthenticateService {
	return &AuthenticateService{ctx: ctx}
}

// Run create note info
func (s *AuthenticateService) Run(req *auth.AuthenticateReq) (resp *auth.AuthenticateResp, err error) {
	// Finish your business logic.

	return
}
