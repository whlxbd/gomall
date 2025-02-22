package service

import (
	"context"
	auth "github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth"
)

type GetService struct {
	ctx context.Context
} // NewGetService new GetService
func NewGetService(ctx context.Context) *GetService {
	return &GetService{ctx: ctx}
}

// Run create note info
func (s *GetService) Run(req *auth.GetReq) (resp *auth.GetResp, err error) {
	// Finish your business logic.

	return
}
