package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	rulemodel "github.com/whlxbd/gomall/app/rule/biz/dal/model/rule"
	"github.com/whlxbd/gomall/app/rule/biz/dal/mysql"
	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth"
	rule "github.com/whlxbd/gomall/rpc_gen/kitex_gen/rule"
	"gorm.io/gorm"

	"github.com/whlxbd/gomall/app/rule/infra/rpc"
)

type CreateService struct {
	ctx context.Context
} // NewCreateService new CreateService
func NewCreateService(ctx context.Context) *CreateService {
	return &CreateService{ctx: ctx}
}

// Run create note info
func (s *CreateService) Run(req *rule.CreateReq) (resp *rule.CreateResp, err error) {
	// Finish your business logic.
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		err = rulemodel.Create(tx, s.ctx, &rulemodel.Rule{
			Role:   req.Role,
			Router: req.Router,
		})
		if err != nil {
			klog.Errorf("create rule failed: %v", err)
			return kerrors.NewBizStatusError(500, "create rule failed")
		}

		_, err = rpc.AuthClient.LoadPolicy(s.ctx, &auth.LoadPolicyReq{
			Role:   req.Role,
			Router: req.Router,
		})
		if err != nil {
			klog.Errorf("load policy failed: %v", err)
			return kerrors.NewBizStatusError(500, "load policy failed")
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	
	return
}
