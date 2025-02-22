package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	rulemodel "github.com/whlxbd/gomall/app/rule/biz/dal/model/rule"
	"github.com/whlxbd/gomall/app/rule/biz/dal/mysql"
	rule "github.com/whlxbd/gomall/rpc_gen/kitex_gen/rule"
)

type ListService struct {
	ctx context.Context
} // NewListService new ListService
func NewListService(ctx context.Context) *ListService {
	return &ListService{ctx: ctx}
}

// Run create note info
func (s *ListService) Run(req *rule.ListReq) (resp *rule.ListResp, err error) {
	// Finish your business logic.
	ruleRows, err := rulemodel.GetPage(mysql.DB, s.ctx, req.Page, req.PageSize)
	if err != nil {
		klog.Errorf("get rule list failed: %v", err)
		return nil, kerrors.NewBizStatusError(400, "get rule list failed")
	}

	klog.Infof("get rule list: %v", ruleRows)

	resp = &rule.ListResp{
		Rules: make([]*rule.Rule, 0, len(ruleRows)),
	}

	for _, ruleRow := range ruleRows {
		resp.Rules = append(resp.Rules, &rule.Rule{
			Id:     int32(ruleRow.ID),
			Role:   ruleRow.Role,
			Router: ruleRow.Router,
		})
	}
	return
}
