package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/whlxbd/gomall/app/rule/biz/dal/cache"
	"github.com/whlxbd/gomall/app/rule/biz/dal/model/whitelist"
	"github.com/whlxbd/gomall/app/rule/biz/dal/mysql"
	"github.com/whlxbd/gomall/app/rule/biz/dal/redis"
	rule "github.com/whlxbd/gomall/rpc_gen/kitex_gen/rule"
	"gorm.io/gorm"
)

type DeleteWhiteRouterService struct {
	ctx context.Context
} // NewDeleteWhiteRouterService new DeleteWhiteRouterService
func NewDeleteWhiteRouterService(ctx context.Context) *DeleteWhiteRouterService {
	return &DeleteWhiteRouterService{ctx: ctx}
}

// Run create note info
func (s *DeleteWhiteRouterService) Run(req *rule.DeleteWhiteRouterReq) (resp *rule.DeleteWhiteRouterResp, err error) {
	// Finish your business logic.
	routerRow, err := whitelist.GetByID(mysql.DB, s.ctx, req.Id)
	if err != nil {
		return nil, err
	}
	err = mysql.DB.Transaction(func(fc *gorm.DB) (err error) {
		err = whitelist.Delete(mysql.DB, s.ctx, req.Id)
		if err != nil {
			klog.Errorf("delete white router failed: %v", err)
			return kerrors.NewBizStatusError(500, "delete white router failed")
		}

		err = cache.DeleteWhiteRouter(s.ctx, redis.RedisClient, routerRow.Router)
		if err != nil {
			klog.Errorf("delete white router failed: %v", err)
			return kerrors.NewBizStatusError(500, "delete white router failed")
		}

		return
	})
	if err != nil {
		return nil, err
	}
	return
}
