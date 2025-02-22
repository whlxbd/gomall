package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/whlxbd/gomall/app/auth/biz/cas"
	auth "github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth"
)

type CheckPermissionService struct {
	ctx context.Context
} // NewCheckPermissionService new CheckPermissionService
func NewCheckPermissionService(ctx context.Context) *CheckPermissionService {
	return &CheckPermissionService{ctx: ctx}
}

// Run create note info
func (s *CheckPermissionService) Run(req *auth.CheckPermissionReq) (resp *auth.CheckPermissionResp, err error) {
	// Finish your business logic.
	err = cas.CheckPolicy(req.Role, req.Router)
	if err != nil {
		klog.Errorf("check permission failed: %v", err)
		return nil, err
	}
	resp = &auth.CheckPermissionResp{
		Ok: true,
	}
	return resp, nil
}
