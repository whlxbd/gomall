package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	rulemodel "github.com/whlxbd/gomall/app/rule/biz/dal/model/rule"
	"github.com/whlxbd/gomall/app/rule/biz/dal/mysql"
	"github.com/whlxbd/gomall/app/rule/infra/rpc"
	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth"
	rule "github.com/whlxbd/gomall/rpc_gen/kitex_gen/rule"
	"gorm.io/gorm"
)

type DeleteService struct {
	ctx context.Context
} // NewDeleteService new DeleteService
func NewDeleteService(ctx context.Context) *DeleteService {
	return &DeleteService{ctx: ctx}
}

// Run create note info
func (s *DeleteService) Run(req *rule.DeleteReq) (resp *rule.DeleteResp, err error) {
	// Finish your business logic.
	ruleRow, err := rulemodel.GetByID(mysql.DB, s.ctx, req.Id)
	if err != nil {
		klog.Errorf("get rule failed: %v", err)
		return nil, kerrors.NewBizStatusError(500, "get rule failed")
	}

	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		err = rulemodel.Delete(tx, s.ctx, req.Id)
		if err != nil {
			klog.Errorf("delete rule failed: %v", err)
			return kerrors.NewBizStatusError(500, "delete rule failed")
		}

		_, err = rpc.AuthClient.RemovePolicy(s.ctx, &auth.RemovePolicyReq{
			Role:   ruleRow.Role,
			Router: ruleRow.Router,
		})
		if err != nil {
			klog.Errorf("remove policy failed: %v", err)
			return kerrors.NewBizStatusError(500, "remove policy failed")
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return
}
