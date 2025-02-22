package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/whlxbd/gomall/app/rule/biz/dal/mysql"
	"github.com/whlxbd/gomall/app/user/biz/dal/model"
	"github.com/whlxbd/gomall/common/utils/authpayload"
	rule "github.com/whlxbd/gomall/rpc_gen/kitex_gen/rule"
)

type DeleteService struct {
	ctx context.Context
} // NewDeleteService new DeleteService
func NewDeleteService(ctx context.Context) *DeleteService {
	return &DeleteService{ctx: ctx}
}

// Run create note info
func (s *DeleteService) Run(req *rule.DeleteReq) (resp *rule.DeleteResp, err error) {
	// Finish your business logic.
	payload, err := authpayload.Get(s.ctx)
	if err != nil {
		klog.Warnf("get auth payload failed: %v", err)
		return nil, kerrors.NewBizStatusError(400, "get auth payload failed")
	}

	if payload.Type != "admin" {
		klog.Warnf("only admin can delete rule")
		return nil, kerrors.NewBizStatusError(400, "only admin can delete rule")
	}

	err = model.Delete(mysql.DB, s.ctx, req.Id)
	if err != nil {
		klog.Errorf("delete rule failed: %v", err)
		return nil, kerrors.NewBizStatusError(500, "delete rule failed")
	}
	return
}
