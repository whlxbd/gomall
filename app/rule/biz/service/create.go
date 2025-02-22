package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/whlxbd/gomall/app/rule/biz/dal/model"
	"github.com/whlxbd/gomall/app/rule/biz/dal/mysql"
	"github.com/whlxbd/gomall/common/utils/authpayload"
	rule "github.com/whlxbd/gomall/rpc_gen/kitex_gen/rule"
)

type CreateService struct {
	ctx context.Context
} // NewCreateService new CreateService
func NewCreateService(ctx context.Context) *CreateService {
	return &CreateService{ctx: ctx}
}

// Run create note info
func (s *CreateService) Run(req *rule.CreateReq) (resp *rule.CreateResp, err error) {
	// Finish your business logic.
	payload, err := authpayload.Get(s.ctx)
	if err != nil {
		klog.Warnf("get auth payload failed: %v", err)
		return nil, kerrors.NewBizStatusError(400, "get auth payload failed")
	}

	if payload.Type != "admin" {
		klog.Warnf("only admin can create note")
		return nil, kerrors.NewBizStatusError(400, "only admin can create rule")
	}

	err = model.Create(mysql.DB, s.ctx, &model.Rule{
		Role:   req.Role,
		Router: req.Router,
	})
	if err != nil {
		klog.Errorf("create rule failed: %v", err)
		return nil, kerrors.NewBizStatusError(500, "create rule failed")
	}
	return
}
