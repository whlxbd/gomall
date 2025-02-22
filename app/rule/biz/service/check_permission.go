package service

import (
	"context"

	"github.com/whlxbd/gomall/app/rule/biz/cas"
	rule "github.com/whlxbd/gomall/rpc_gen/kitex_gen/rule"
)

type CheckPermissionService struct {
	ctx context.Context
} // NewCheckPermissionService new CheckPermissionService
func NewCheckPermissionService(ctx context.Context) *CheckPermissionService {
	return &CheckPermissionService{ctx: ctx}
}

// Run create note info
func (s *CheckPermissionService) Run(req *rule.CheckPermissionReq) (resp *rule.CheckPermissionResp, err error) {
	// Finish your business logic.
	err = cas.CheckPolicy(req.Role, req.Router)
	if err != nil {
		return nil, err
	}
	return
}
