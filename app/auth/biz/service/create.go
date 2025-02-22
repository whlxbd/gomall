package service

import (
	"context"
	auth "github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth"
)

type CreateService struct {
	ctx context.Context
} // NewCreateService new CreateService
func NewCreateService(ctx context.Context) *CreateService {
	return &CreateService{ctx: ctx}
}

// Run create note info
func (s *CreateService) Run(req *auth.CreateReq) (resp *auth.CreateResp, err error) {
	// Finish your business logic.

	return
}
