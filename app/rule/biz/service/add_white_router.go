package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/whlxbd/gomall/app/rule/biz/dal/cache"
	"github.com/whlxbd/gomall/app/rule/biz/dal/model/whitelist"
	"github.com/whlxbd/gomall/app/rule/biz/dal/mysql"
	rule "github.com/whlxbd/gomall/rpc_gen/kitex_gen/rule"
	"gorm.io/gorm"
)

type AddWhiteRouterService struct {
	ctx context.Context
} // NewAddWhiteRouterService new AddWhiteRouterService
func NewAddWhiteRouterService(ctx context.Context) *AddWhiteRouterService {
	return &AddWhiteRouterService{ctx: ctx}
}

// Run create note info
func (s *AddWhiteRouterService) Run(req *rule.AddWhiteRouterReq) (resp *rule.AddWhiteRouterResp, err error) {
	// Finish your business logic.
	err = mysql.DB.Transaction(func(fc *gorm.DB) (err error) {
		err = whitelist.Create(mysql.DB, s.ctx, &whitelist.WhiteRouter{
			Router: req.Router,
		})
		if err != nil {
			klog.Errorf("create white router failed: %v", err)
			return kerrors.NewBizStatusError(500, "create white router failed")
		}

		err = cache.SetWhiteRouter(s.ctx, req.Router)
		if err != nil {
			klog.Errorf("set white router failed: %v", err)
			return kerrors.NewBizStatusError(500, "set white router failed")
		}
		return err
	})
	if err != nil {
		return nil, err
	}

	return
}
