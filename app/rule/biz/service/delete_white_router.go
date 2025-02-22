package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/whlxbd/gomall/app/rule/biz/dal/model/whitelist"
	"github.com/whlxbd/gomall/app/rule/biz/dal/mysql"
	rule "github.com/whlxbd/gomall/rpc_gen/kitex_gen/rule"
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
	err = whitelist.Delete(mysql.DB, s.ctx, req.Id)
	if err != nil {
		klog.Errorf("delete white router failed: %v", err)
		return nil, kerrors.NewBizStatusError(500, "delete white router failed")
	}
	return
}
