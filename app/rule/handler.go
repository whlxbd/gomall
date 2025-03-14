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

// GetWhiteList implements the RuleServiceImpl interface.
func (s *RuleServiceImpl) GetWhiteList(ctx context.Context, req *rule.GetWhiteListReq) (resp *rule.GetWhiteListResp, err error) {
	resp, err = service.NewGetWhiteListService(ctx).Run(req)

	return resp, err
}

// AddWhiteRouter implements the RuleServiceImpl interface.
func (s *RuleServiceImpl) AddWhiteRouter(ctx context.Context, req *rule.AddWhiteRouterReq) (resp *rule.AddWhiteRouterResp, err error) {
	resp, err = service.NewAddWhiteRouterService(ctx).Run(req)

	return resp, err
}

// DeleteWhiteRouter implements the RuleServiceImpl interface.
func (s *RuleServiceImpl) DeleteWhiteRouter(ctx context.Context, req *rule.DeleteWhiteRouterReq) (resp *rule.DeleteWhiteRouterResp, err error) {
	resp, err = service.NewDeleteWhiteRouterService(ctx).Run(req)

	return resp, err
}
