package service

import (
	"context"
	"os"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/golang-jwt/jwt/v5"
	auth "github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth"
)

type VerifyTokenByRPCService struct {
	ctx context.Context
} // NewVerifyTokenByRPCService new VerifyTokenByRPCService
func NewVerifyTokenByRPCService(ctx context.Context) *VerifyTokenByRPCService {
	return &VerifyTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *VerifyTokenByRPCService) Run(req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	// Finish your business logic.
	token, err := jwt.ParseWithClaims(req.Token, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		klog.Errorf("parse token failed: %v", err)
		return nil, kerrors.NewBizStatusError(400, "parse token failed")
	}
	resp = &auth.VerifyResp{
		Res: token.Valid,
	}
	return
}
