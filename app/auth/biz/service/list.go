package service

import (
	"context"
	auth "github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth"
)

type ListService struct {
	ctx context.Context
} // NewListService new ListService
func NewListService(ctx context.Context) *ListService {
	return &ListService{ctx: ctx}
}

// Run create note info
func (s *ListService) Run(req *auth.ListReq) (resp *auth.ListResp, err error) {
	// Finish your business logic.

	return
}
