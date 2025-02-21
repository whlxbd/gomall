package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/go-playground/validator/v10"
	"github.com/whlxbd/gomall/app/user/biz/dal/model"
	"github.com/whlxbd/gomall/app/user/biz/dal/mysql"
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
		klog.Error("email is empty")
		return nil, kerrors.NewBizStatusError(400, "email is empty")
	}
	validate := validator.New()
	err = validate.Var(req.Email, "required,email")
	if err != nil {
		klog.Error("invalid email format")
		return nil, kerrors.NewBizStatusError(400, "invalid email format")
	}
	if req.Password == "" {
		klog.Error("password is empty")
		return nil, kerrors.NewBizStatusError(400, "password is empty")
	}
	if req.ConfirmPassword == "" {
		klog.Error("ConfirmPassword is empty")
		return nil, kerrors.NewBizStatusError(400, "ConfirmPassword is empty")
	}
	if req.Password != req.ConfirmPassword {
		klog.Error("password and ConfirmPassword are not the same")
		return nil, kerrors.NewBizStatusError(400, "password and ConfirmPassword are not the same")
	}
	if findUser, err := model.GetByEmail(mysql.DB, context.Background(), req.Email); err == nil && findUser != nil {
		klog.Error("user already exists")
		return nil, kerrors.NewBizStatusError(400, "user already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	err = model.Create(mysql.DB, context.Background(), &model.User{
		Email:    req.Email,
		Password: string(hashedPassword),
	})
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewGRPCBizStatusError(500, "create user failed")
	}
	userRow, err := model.GetByEmail(mysql.DB, context.Background(), req.Email)
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(500, "get user info failed")
	}
	resp = &user.RegisterResp{
		UserId: userRow.ID,
	}
	return
}
