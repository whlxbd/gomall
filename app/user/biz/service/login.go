package service

import (
	"context"
	"errors"
	"github.com/whlxbd/gomall/app/user/biz/dal/mysql"
	"github.com/whlxbd/gomall/app/user/biz/model"
	user "github.com/whlxbd/gomall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.
	userRow, err := model.GetByEmail(mysql.DB, context.Background(), req.Email)
	if err != nil {
		return nil, err
	}
	if userRow == nil {
		return nil, errors.New("not found user")
	}
	if err = bcrypt.CompareHashAndPassword([]byte(userRow.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("password not match")
	}
	resp = &user.LoginResp{
		UserId: int32(userRow.ID),
	}

	return
}
