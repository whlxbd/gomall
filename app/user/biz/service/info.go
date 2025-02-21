package service

import (
	"context"
	"encoding/json"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	cache "github.com/whlxbd/gomall/app/user/biz/dal/cache"
	"github.com/whlxbd/gomall/app/user/biz/dal/model"
	"github.com/whlxbd/gomall/app/user/biz/dal/mysql"
	user "github.com/whlxbd/gomall/rpc_gen/kitex_gen/user"

	goredis "github.com/redis/go-redis/v9"
	userredis "github.com/whlxbd/gomall/app/user/biz/dal/redis"
)

type InfoService struct {
	ctx context.Context
} // NewInfoService new InfoService
func NewInfoService(ctx context.Context) *InfoService {
	return &InfoService{ctx: ctx}
}

// Run create note info
func (s *InfoService) Run(req *user.InfoReq) (resp *user.InfoResp, err error) {
	// Finish your business logic.
	userCache, err := cache.GetByID(userredis.RedisClient, s.ctx, req.UserId)
	if err != nil && err != goredis.Nil {
		klog.Errorf("get user info from cache failed: %v", err)
		return nil, kerrors.NewBizStatusError(400, "redis get failed")
	}
	if err == nil {
		err = json.Unmarshal([]byte(userCache), &resp)
		if err != nil {
			klog.Errorf("unmarshal user info failed: %v", err)
			return nil, kerrors.NewBizStatusError(400, "unmarshal user info failed")
		}
		return
	}

	userRow, err := model.GetByID(mysql.DB, s.ctx, req.UserId)
	if err != nil {
		klog.Errorf("get user info failed: %v", err)
		return nil, kerrors.NewBizStatusError(400, "user not found")
	}
	resp = &user.InfoResp{
		UserId:    userRow.ID,
		Email:     userRow.Email,
		Username:  userRow.UserName,
		AvatarUrl: userRow.AvatarUrl,
		Type:      userRow.Type,
	}

	tmp, err := json.Marshal(resp)
	if err != nil {
		klog.Errorf("marshal user info failed: %v", err)
		return nil, kerrors.NewBizStatusError(400, "marshal user info failed")
	}
	err = cache.SetByID(userredis.RedisClient, s.ctx, req.UserId, string(tmp))
	if err != nil {
		klog.Errorf("set user info to cache failed: %v", err)
		return nil, kerrors.NewBizStatusError(400, "redis set failed")
	}

	return
}
