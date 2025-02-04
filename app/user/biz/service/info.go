package service

import (
	"context"
	"errors"

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
	userRow, err := model.GetByID(mysql.DB, s.ctx, int(req.UserId))
	if err != nil {
		return nil, errors.New("not found user")
	}
	resp = &user.InfoResp{
		Email: userRow.Email,
	}

	return
}
