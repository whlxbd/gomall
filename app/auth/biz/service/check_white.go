package service

import (
	"context"

	"github.com/whlxbd/gomall/app/auth/biz/dal/redis"
	auth "github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth"

	rulecache "github.com/whlxbd/gomall/app/rule/biz/dal/cache"
)

type CheckWhiteService struct {
	ctx context.Context
} // NewCheckWhiteService new CheckWhiteService
func NewCheckWhiteService(ctx context.Context) *CheckWhiteService {
	return &CheckWhiteService{ctx: ctx}
}

// Run create note info
func (s *CheckWhiteService) Run(req *auth.CheckWhiteReq) (resp *auth.CheckWhiteResp, err error) {
	// Finish your business logic.
	whiteRouterRow, err := rulecache.GetWhiteRouter(s.ctx, redis.RedisClient, req.Router)
	
	if err != nil {
		return nil, err
	}

	if whiteRouterRow == nil || whiteRouterRow.Router != req.Router {
		return &auth.CheckWhiteResp{Ok: false}, nil
	} else {
		return &auth.CheckWhiteResp{Ok: true}, nil
	}
}
