package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/whlxbd/gomall/common/utils/authpayload"
	rule "github.com/whlxbd/gomall/rpc_gen/kitex_gen/rule"
)

type UpdateService struct {
	ctx context.Context
} // NewUpdateService new UpdateService
func NewUpdateService(ctx context.Context) *UpdateService {
	return &UpdateService{ctx: ctx}
}

// Run create note info
func (s *UpdateService) Run(req *rule.UpdateReq) (resp *rule.UpdateResp, err error) {
	// Finish your business logic.
	payload, err := authpayload.Get(s.ctx)
	if err != nil {
		klog.Warnf("get auth payload failed: %v", err)
		return nil, kerrors.NewBizStatusError(400, "get auth payload failed")
	}

	if payload.Type != "admin" {
		klog.Warnf("only admin can update rule")
		return nil, kerrors.NewBizStatusError(400, "only admin can update rule")
	}

	return nil, kerrors.NewBizStatusError(400, "update rule does not support")

	// rule := &model.Rule{
	// 	Role:   req.Rule.Role,
	// 	Router: req.Rule.Router,
	// }

	// rule.ID = uint(req.Rule.Id)

	// err = model.Update(mysql.DB, s.ctx, rule)
	// return
}
