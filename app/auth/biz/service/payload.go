package service

import (
	"context"
	"os"

	"github.com/golang-jwt/jwt/v5"
	auth "github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth"
)

type PayloadService struct {
	ctx context.Context
} // NewPayloadService new PayloadService
func NewPayloadService(ctx context.Context) *PayloadService {
	return &PayloadService{ctx: ctx}
}

// Run create note info
func (s *PayloadService) Run(req *auth.PayloadReq) (resp *auth.PayloadResp, err error) {
	// Finish your business logic.
	payload := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(req.Token, payload, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return &auth.PayloadResp{
		UserId: int32((payload["user_id"].(float64))),
		Type:   payload["type"].(string),
	}, nil
}
