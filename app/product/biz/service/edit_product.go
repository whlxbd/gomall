package service

import (
	"context"
	product "github.com/whlxbd/gomall/rpc_gen/kitex_gen/product"
	"github.com/whlxbd/gomall/app/product/biz/dal/mysql"
	"github.com/whlxbd/gomall/app/product/biz/dal/redis"
	"github.com/whlxbd/gomall/app/product/biz/dal/model"
	"github.com/whlxbd/gomall/common/utils/pool"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type EditProductService struct {
	ctx context.Context
} // NewEditProductService new EditProductService
func NewEditProductService(ctx context.Context) *EditProductService {
	return &EditProductService{ctx: ctx}
}

// Run create note info
func (s *EditProductService) Run(req *product.EditProductReq) (resp *product.EditProductResp, err error) {
	// Finish your business logic.
	// if middleware.CheckAdminPermission(s.ctx) != nil {
	// 	return nil, kerrors.NewBizStatusError(400, "permission denied: admin required")
	// }
	_ = pool.Submit(func() {
		klog.Infof("edit product: %+v", req)
	})
	
	resp = &product.EditProductResp{}


	if req.Id == 0 { // 商品ID不能为空
		_ = pool.Submit(func() {
			klog.Errorf("id is required: %+v", req)
		})
		return nil, kerrors.NewBizStatusError(400, "id is required")
	}

	if req.Name == "" { // 商品名称不能为空
		_ = pool.Submit(func() {
			klog.Errorf("name is required: %+v", req)
		})
		return nil, kerrors.NewBizStatusError(400, "name is required")
	}

	if req.Price < 0 || req.Stock < 0 || req.Soldcount < 0 { // 价格、库存、销量必须大于0
		_ = pool.Submit(func() {
			klog.Errorf("price, stock, soldcount, must be greater than 0: %+v", req)
		})
		return nil, kerrors.NewBizStatusError(400, "price, stock, soldcount, must be greater than 0")
	}

	exist, err := model.NewCachedProductQuery(model.NewProductQuery(s.ctx, mysql.DB), redis.RedisClient).GetById(req.Id)

	if err != nil { // 查询商品失败
		resp.Success = false
		return resp, kerrors.NewBizStatusError(400, err.Error())
	}

	if exist == nil { // 商品不存在
		klog.Errorf("product not found, id: %d\n", req.Id)
		resp.Success = false
		return resp, kerrors.NewBizStatusError(400, "product not found")
	}

	if err := model.ProductStatus(req.Status).IsValid(); err != nil { // 商品状态不合法
		resp.Success = false
		return resp, kerrors.NewBizStatusError(400, "status is invalid")
	}

	categories, err := String2Category(&req.Categories, &s.ctx)

	if err != nil { 
		resp.Success = false
		return
	}

	p := &model.Product{
		Base: model.Base{
			ID: req.Id,
		},
		Name:        req.Name,
		Description: req.Description,
		Picture:     req.Picture,
		Price:       req.Price,
		Stock:       req.Stock,
		SoldCount:   req.Soldcount,
		Status:      model.ProductStatus(req.Status),
		IsHot:       req.Ishot,
		IsNew:       req.Isnew,
		IsRecommend: req.Isrecommend,
		Categories:  *categories,
	}

	if err = model.EditProduct(mysql.DB, redis.RedisClient, s.ctx, p); err != nil { // 更新商品
		klog.Errorf("update product failed, id: %d, err: %v\n", req.Id, err)
		return nil, kerrors.NewBizStatusError(400, err.Error())
	}

	resp.Success = true
	return
}
