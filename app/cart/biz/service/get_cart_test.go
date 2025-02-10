package service

import (
	"context"
	"testing"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/stretchr/testify/assert"
	"github.com/whlxbd/gomall/app/cart/biz/dal/model"
	"github.com/whlxbd/gomall/app/cart/biz/dal/mysql"
	cart "github.com/whlxbd/gomall/rpc_gen/kitex_gen/cart"
)

func TestGetCart_Run(t *testing.T) {
	setupTestDB(t)

	// 准备测试数据
	testCarts := []*model.Cart{
		{
			UserID:    1,
			ProductID: 1,
			Quantity:  2,
			Selected:  true,
			Status:    true,
		},
		{
			UserID:    1,
			ProductID: 2,
			Quantity:  3,
			Selected:  true,
			Status:    true,
		},
	}

	// 写入测试数据
	for _, cart := range testCarts {
		err := mysql.DB.Create(cart).Error
		assert.NoError(t, err)
	}

	tests := []struct {
		name    string
		req     *cart.GetCartReq
		want    *cart.GetCartResp
		wantErr bool
		errCode int32
	}{
		{
			name: "获取购物车成功",
			req: &cart.GetCartReq{
				UserId: 1,
			},
			want: &cart.GetCartResp{
				Cart: &cart.Cart{
					UserId: 1,
					Items: []*cart.CartItem{
						{ProductId: 1, Quantity: 2},
						{ProductId: 2, Quantity: 3},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "购物车为空",
			req: &cart.GetCartReq{
				UserId: 999,
			},
			wantErr: true,
			errCode: 40007,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewGetCartService(context.Background())
			resp, err := s.Run(tt.req)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetCartService.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				if bizErr, ok := kerrors.FromBizStatusError(err); ok && bizErr.BizStatusCode() != tt.errCode {
					t.Errorf("GetCartService.Run() error code = %v, wantErr %v", bizErr.BizStatusCode(), tt.errCode)
				}
				return
			}

			assert.Equal(t, tt.want.Cart.UserId, resp.Cart.UserId)
			assert.Equal(t, len(tt.want.Cart.Items), len(resp.Cart.Items))

			for i, item := range resp.Cart.Items {
				assert.Equal(t, tt.want.Cart.Items[i].ProductId, item.ProductId)
				assert.Equal(t, tt.want.Cart.Items[i].Quantity, item.Quantity)
			}
		})
	}
}
