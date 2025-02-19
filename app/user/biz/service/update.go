package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/whlxbd/gomall/app/user/biz/dal/mysql"
	"github.com/whlxbd/gomall/app/user/biz/model"
	"github.com/whlxbd/gomall/common/utils/authpayload"
	user "github.com/whlxbd/gomall/rpc_gen/kitex_gen/user"
)

type UpdateService struct {
	ctx context.Context
} // NewUpdateService new UpdateService
func NewUpdateService(ctx context.Context) *UpdateService {
	return &UpdateService{ctx: ctx}
}

// Run create note info
func (s *UpdateService) Run(req *user.UpdateReq) (resp *user.UpdateResp, err error) {
	// Finish your business logic.
	payload, err := authpayload.Get(s.ctx)
	if err != nil {
		return nil, kerrors.NewBizStatusError(400, "get payload failed")
	}
	if payload.Type != "admin" || payload.UserId != req.UserId {
		return nil, kerrors.NewBizStatusError(400, "permission denied")
	}
	resp = &user.UpdateResp{
		Success: false,
	}
	userRow, err := model.GetByID(mysql.DB, s.ctx, req.UserId)
	if err != nil {
		return resp, kerrors.NewBizStatusError(400, "user not found")
	}
	userRow.Email = req.Email
	userRow.UserName = req.Username
	userRow.AvatarUrl = req.AvatarUrl
	userRow.Type = req.Type
	err = model.Update(mysql.DB, s.ctx, userRow)
	if err != nil {
		return resp, kerrors.NewBizStatusError(400, "update user failed")
	}
	resp.Success = true
	return
}
