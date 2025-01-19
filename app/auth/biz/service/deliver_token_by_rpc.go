package service

import (
	"context"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
		str   string
	)

	key = []byte(os.Getenv("JWT_SECRET"))
	t = jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": req.UserId,
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
