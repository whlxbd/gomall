package service

import (
	"context"
	"os"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/golang-jwt/jwt/v5"
	auth "github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth"
)

type GetPayloadService struct {
	ctx context.Context
} // NewGetPayloadService new GetPayloadService
func NewGetPayloadService(ctx context.Context) *GetPayloadService {
	return &GetPayloadService{ctx: ctx}
}

// Run create note info
func (s *GetPayloadService) Run(req *auth.GetPayloadReq) (resp *auth.GetPayloadResp, err error) {
	// Finish your business logic.
	payload := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(req.Token, payload, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		klog.Errorf("parse token failed: %v", err)
		return nil, kerrors.NewBizStatusError(400, "parse token failed")
	}
	return &auth.GetPayloadResp{
		UserId: int32((payload["user_id"].(float64))),
		Type:   payload["type"].(string),
	}, nil
}
