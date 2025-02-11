package service

import (
	"context"
	"errors"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/go-playground/validator/v10"
	"github.com/whlxbd/gomall/app/user/biz/dal/mysql"
	"github.com/whlxbd/gomall/app/user/biz/model"
	user "github.com/whlxbd/gomall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// Finish your business logic.
	if req.Email == "" {
		return nil, errors.New("email is empty")
	}
	validate := validator.New()
	err = validate.Var(req.Email, "required,email")
	if err != nil {
		return nil, errors.New("invalid email format")
	}
	if req.Password == "" {
		return nil, errors.New("password is empty")
	}
	if req.ConfirmPassword == "" {
		return nil, errors.New("ConfirmPassword is empty")
	}
	if req.Password != req.ConfirmPassword {
		return nil, errors.New("password and ConfirmPassword are not the same")
	}
	if findUser, err := model.GetByEmail(mysql.DB, context.Background(), req.Email); err == nil && findUser != nil {
		return nil, errors.New("user already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	err = model.Create(mysql.DB, context.Background(), &model.User{
		Email:    req.Email,
		Password: string(hashedPassword),
	})
	if err != nil {
		klog.Error(err)
	}
	userRow, err := model.GetByEmail(mysql.DB, context.Background(), req.Email)
	if err != nil {
		return nil, errors.New("not found user")
	}
	resp = &user.RegisterResp{
		UserId: userRow.ID,
	}
	return
}
