package service

import (
	"context"
	"os"
	"time"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/golang-jwt/jwt/v5"
	"github.com/whlxbd/gomall/app/auth/infra/rpc"
	user "github.com/whlxbd/gomall/rpc_gen/kitex_gen/user"
	auth "github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth"
)

type DeliverTokenByRPCService struct {
	ctx context.Context
} // NewDeliverTokenByRPCService new DeliverTokenByRPCService
func NewDeliverTokenByRPCService(ctx context.Context) *DeliverTokenByRPCService {
	return &DeliverTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *DeliverTokenByRPCService) Run(req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	// Finish your business logic.

	var (
		key []byte
		t   *jwt.Token
		str string
	)

	key = []byte(os.Getenv("JWT_SECRET"))
	userInfo, err := rpc.UserClient.Info(s.ctx, &user.InfoReq{UserId: req.UserId})
	if err != nil {
		klog.Errorf("get user info failed: %v", err)
		return nil, kerrors.NewBizStatusError(400, "user not found")
	}

	t = jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id":     req.UserId,
			"type":        userInfo.Type,
			"expire_time": time.Now().Add(time.Hour * 24).Unix(),
		})
	str, err = t.SignedString(key)
	if err != nil {
		klog.Errorf("sign token failed: %v", err)
		return nil, kerrors.NewBizStatusError(400, "sign token failed")
	}

	resp = &auth.DeliveryResp{
		Token: str,
	}
	return
}
