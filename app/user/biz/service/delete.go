package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/whlxbd/gomall/app/user/biz/dal/model"
	"github.com/whlxbd/gomall/app/user/biz/dal/mysql"
	"github.com/whlxbd/gomall/common/utils/authpayload"
	user "github.com/whlxbd/gomall/rpc_gen/kitex_gen/user"
)

type DeleteService struct {
	ctx context.Context
} // NewDeleteService new DeleteService
func NewDeleteService(ctx context.Context) *DeleteService {
	return &DeleteService{ctx: ctx}
}

// Run create note info
func (s *DeleteService) Run(req *user.DeleteReq) (resp *user.DeleteResp, err error) {
	// Finish your business logic.
	payload, err := authpayload.Get(s.ctx)
	if err != nil {
		klog.Errorf("get payload failed: %v", err)
		return nil, kerrors.NewBizStatusError(400, "get payload failed")
	}
	if payload.Type != "admin" {
		return nil, kerrors.NewBizStatusError(400, "permission denied")
	}
	resp = &user.DeleteResp{
		Success: false,
	}
	_, err = model.GetByID(mysql.DB, s.ctx, req.UserId)
	if err != nil {
		klog.Errorf("get user info failed: %v", err)
		return resp, kerrors.NewBizStatusError(400, "user not found")
	}
	err = model.Delete(mysql.DB, s.ctx, req.UserId)
	if err != nil {
		klog.Errorf("delete user failed: %v", err)
		return resp, kerrors.NewBizStatusError(400, "delete user failed")
	}
	resp = &user.DeleteResp{
		Success: true,
	}
	return
}
