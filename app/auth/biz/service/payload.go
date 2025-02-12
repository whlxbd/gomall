package service

import (
	"context"
	auth "github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth"
)

type PayloadService struct {
	ctx context.Context
} // NewPayloadService new PayloadService
func NewPayloadService(ctx context.Context) *PayloadService {
	return &PayloadService{ctx: ctx}
}

// Run create note info
func (s *PayloadService) Run(req *auth.PayloadReq) (resp *auth.PayloadResp, err error) {
	// Finish your business logic.

	return
}
