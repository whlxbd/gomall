package service

import (
	"context"
	auth "github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth"
)

type LoadPolicyService struct {
	ctx context.Context
} // NewLoadPolicyService new LoadPolicyService
func NewLoadPolicyService(ctx context.Context) *LoadPolicyService {
	return &LoadPolicyService{ctx: ctx}
}

// Run create note info
func (s *LoadPolicyService) Run(req *auth.LoadPolicyReq) (resp *auth.LoadPolicyResp, err error) {
	// Finish your business logic.

	return
}
