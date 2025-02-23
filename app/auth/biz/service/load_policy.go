package service

import (
	"context"

	"github.com/whlxbd/gomall/app/auth/biz/cas"
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
	err = cas.AddPolicy(req.Role, req.Router)
	if err != nil {
		return nil, err
	}

	return
}
