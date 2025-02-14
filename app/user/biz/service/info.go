package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/whlxbd/gomall/app/user/biz/dal/mysql"
	"github.com/whlxbd/gomall/app/user/biz/model"
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
	userRow, err := model.GetByID(mysql.DB, s.ctx, req.UserId)
	if err != nil {
		klog.Errorf("get user info failed: %v", err)
		return nil, kerrors.NewBizStatusError(400, "user not found")
	}
	resp = &user.InfoResp{
		UserId:    userRow.ID,
		Email:     userRow.Email,
		Username:  userRow.UserName,
		AvatarUrl: userRow.AvatarUrl,
		Type:      userRow.Type,
	}

	return
}
