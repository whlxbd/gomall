package service

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/whlxbd/gomall/app/user/biz/dal/mysql"
	"github.com/whlxbd/gomall/app/user/biz/model"
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
	userRow, err := model.GetByID(mysql.DB, s.ctx, req.UserId)
	if err != nil {
		return nil, errors.New("user not found")
	}

	t = jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id":     req.UserId,
			"type":        userRow.Type,
			"expire_time": time.Now().Add(time.Hour * 24).Unix(),
		})
	str, err = t.SignedString(key)
	if err != nil {
		return nil, err
	}

	resp = &auth.DeliveryResp{
		Token: str,
	}
	return
}
