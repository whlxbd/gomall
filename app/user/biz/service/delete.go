package service

import (
	"context"
	"errors"

	"github.com/whlxbd/gomall/app/user/biz/dal/mysql"
	"github.com/whlxbd/gomall/app/user/biz/model"
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
	resp = &user.DeleteResp{
		Success: false,
	}
	_, err = model.GetByID(mysql.DB, s.ctx, req.UserId)
	if err != nil {
		return resp, errors.New("user not found")
	}
	err = model.Delete(mysql.DB, s.ctx, req.UserId)
	if err != nil {
		return resp, errors.New("delete user failed")
	}
	resp = &user.DeleteResp{
		Success: true,
	}
	return
}
