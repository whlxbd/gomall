package service

import (
	"context"
	auth "github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth"
)

type RemovePolicyService struct {
	ctx context.Context
} // NewRemovePolicyService new RemovePolicyService
func NewRemovePolicyService(ctx context.Context) *RemovePolicyService {
	return &RemovePolicyService{ctx: ctx}
}

// Run create note info
func (s *RemovePolicyService) Run(req *auth.RemovePolicyReq) (resp *auth.RemovePolicyResp, err error) {
	// Finish your business logic.

	return
}
