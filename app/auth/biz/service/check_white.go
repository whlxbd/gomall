package service

import (
	"context"
	auth "github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth"
	"github.com/rule/"
)

type CheckWhiteService struct {
	ctx context.Context
} // NewCheckWhiteService new CheckWhiteService
func NewCheckWhiteService(ctx context.Context) *CheckWhiteService {
	return &CheckWhiteService{ctx: ctx}
}

// Run create note info
func (s *CheckWhiteService) Run(req *auth.CheckWhiteReq) (resp *auth.CheckWhiteResp, err error) {
	// Finish your business logic.

	return
}
