package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/whlxbd/gomall/app/rule/biz/dal/model/whitelist"
	"github.com/whlxbd/gomall/app/rule/biz/dal/mysql"
	rule "github.com/whlxbd/gomall/rpc_gen/kitex_gen/rule"
)

type GetWhiteListService struct {
	ctx context.Context
} // NewGetWhiteListService new GetWhiteListService
func NewGetWhiteListService(ctx context.Context) *GetWhiteListService {
	return &GetWhiteListService{ctx: ctx}
}

// Run create note info
func (s *GetWhiteListService) Run(req *rule.GetWhiteListReq) (resp *rule.GetWhiteListResp, err error) {
	// Finish your business logic.
	whiteRouterRows, err := whitelist.GetPage(mysql.DB, s.ctx, req.Page, req.PageSize)
	if err != nil {
		klog.Errorf("get white router failed: %v", err)
		return nil, kerrors.NewBizStatusError(500, "get white router failed")
	}

	resp = &rule.GetWhiteListResp{
		List: make([]*rule.WhiteRouter, 0, len(whiteRouterRows)),
	}

	for _, whiteRouterRow := range whiteRouterRows {
		resp.List = append(resp.List, &rule.WhiteRouter{
			Id:     int32(whiteRouterRow.ID),
			Router: whiteRouterRow.Router,
		})
	}

	return
}
