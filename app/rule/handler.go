package main

import (
	"context"
	rule "github.com/whlxbd/gomall/rpc_gen/kitex_gen/rule"
	"github.com/whlxbd/gomall/app/rule/biz/service"
)

// RuleServiceImpl implements the last service interface defined in the IDL.
type RuleServiceImpl struct{}

// Create implements the RuleServiceImpl interface.
func (s *RuleServiceImpl) Create(ctx context.Context, req *rule.CreateReq) (resp *rule.CreateResp, err error) {
	resp, err = service.NewCreateService(ctx).Run(req)

	return resp, err
}

// List implements the RuleServiceImpl interface.
func (s *RuleServiceImpl) List(ctx context.Context, req *rule.ListReq) (resp *rule.ListResp, err error) {
	resp, err = service.NewListService(ctx).Run(req)

	return resp, err
}

// Delete implements the RuleServiceImpl interface.
func (s *RuleServiceImpl) Delete(ctx context.Context, req *rule.DeleteReq) (resp *rule.DeleteResp, err error) {
	resp, err = service.NewDeleteService(ctx).Run(req)

	return resp, err
}

// Get implements the RuleServiceImpl interface.
func (s *RuleServiceImpl) Get(ctx context.Context, req *rule.GetReq) (resp *rule.GetResp, err error) {
	resp, err = service.NewGetService(ctx).Run(req)

	return resp, err
}

// Update implements the RuleServiceImpl interface.
func (s *RuleServiceImpl) Update(ctx context.Context, req *rule.UpdateReq) (resp *rule.UpdateResp, err error) {
	resp, err = service.NewUpdateService(ctx).Run(req)

	return resp, err
}

// CheckPermission implements the RuleServiceImpl interface.
func (s *RuleServiceImpl) CheckPermission(ctx context.Context, req *rule.CheckPermissionReq) (resp *rule.CheckPermissionResp, err error) {
	resp, err = service.NewCheckPermissionService(ctx).Run(req)

	return resp, err
}
