package service

import (
	"context"

	rulemodel "github.com/whlxbd/gomall/app/rule/biz/dal/model/rule"
	"github.com/whlxbd/gomall/app/rule/biz/dal/mysql"
	rule "github.com/whlxbd/gomall/rpc_gen/kitex_gen/rule"
)

type GetService struct {
	ctx context.Context
} // NewGetService new GetService
func NewGetService(ctx context.Context) *GetService {
	return &GetService{ctx: ctx}
}

// Run create note info
func (s *GetService) Run(req *rule.GetReq) (resp *rule.GetResp, err error) {
	// Finish your business logic.
	ruleRow, err := rulemodel.GetByID(mysql.DB, s.ctx, req.Id)
	if err != nil {
		return nil, err
	}

	resp = &rule.GetResp{
		Rule: &rule.Rule{
			Id:     int32(ruleRow.ID),
			Role:   ruleRow.Role,
			Router: ruleRow.Router,
		},
	}
	return
}
