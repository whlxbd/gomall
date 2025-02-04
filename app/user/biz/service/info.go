package service

import (
	"context"
	user "github.com/whlxbd/gomall/rpc_gen/kitex_gen/user"
)

type InfoService struct {
	ctx context.Context
} // NewInfoService new InfoService
func NewInfoService(ctx context.Context) *InfoService {
	return &InfoService{ctx: ctx}
}

// Run create note info
func (s *InfoService) Run(req *user.InfoReq) (resp *user.InfoResp, err error) {
	// Finish your business logic.

	return
}
