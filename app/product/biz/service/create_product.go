package service

import (
	"context"
	// "fmt"

	"github.com/whlxbd/gomall/app/product/biz/dal/model"
	"github.com/whlxbd/gomall/app/product/biz/dal/mysql"
	"github.com/whlxbd/gomall/app/product/biz/dal/redis"

	"github.com/whlxbd/gomall/common/utils/pool"
	// "github.com/whlxbd/gomall/app/product/biz/middleware"
	product "github.com/whlxbd/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"

	"github.com/cloudwego/kitex/pkg/klog"
)

type CreateProductService struct {
	ctx context.Context
} // NewCreateProductService new CreateProductService
func NewCreateProductService(ctx context.Context) *CreateProductService {
	return &CreateProductService{ctx: ctx}
}

// 将分类名称转换为分类对象，首先查询缓存，如果缓存中不存在则查询数据库
func String2Category(ReqCategories *[]string, ctx *context.Context) (*[]model.Category, error) {
	ccq := model.NewCachedCategoryQuery(model.NewCategoryQuery(*ctx, mysql.DB), redis.RedisClient)
	categories := make([]model.Category, len(*ReqCategories))

	for i, cat := range *ReqCategories {
		categories[i] = model.Category{Name: cat}

		// 先查询缓存
		id, err := ccq.IsExistByName(cat)
		if err != nil {
			klog.Error("查询分类失败", err)
			return nil, kerrors.NewBizStatusError(400, "查询分类失败")
		}

		if id > 0 {
			categories[i].ID = id
			continue
		}

		// 查找或创建分类
		if err := mysql.DB.Where("name = ?", cat).
			FirstOrCreate(&categories[i], model.Category{Name: cat}).Error; err != nil {
			klog.Error("创建分类失败", err)
			return nil, kerrors.NewBizStatusError(400, "创建分类失败")
		}

		// 异步更新缓存
		_ = pool.Submit(func() {
			if err := ccq.SetCache(categories[i].Name, categories[i].ID); err != nil {
				klog.Error("设置分类缓存失败", err)
			}
		})

	}
	return &categories, nil
}

// Run create note info
func (s *CreateProductService) Run(req *product.CreateProductReq) (resp *product.CreateProductResp, err error) {
	// Finish your business logic.
	// klog.Infof("CreateProductService input: %+v", req)
	// fmt.Println("CreateProductService input: ", req)
	// fmt.Println("s.ctx: ", s.ctx)

	// 无法传入metadata
	// if err = middleware.CheckAdminPermission(s.ctx); err != nil {
	// 	return nil, kerrors.NewBizStatusError(400, err.Error())
	// }

	klog.Infof("CreateProductService input: %+v\n", req)

	if req.Name == "" {
		return nil, kerrors.NewBizStatusError(400, "name is required")
	}

	if req.Price < 0 || req.Stock < 0 || req.Soldcount < 0 {
		return nil, kerrors.NewBizStatusError(400, "price, stock, soldcount, must be greater than 0")
	}

	if err := model.ProductStatus(req.Status).IsValid(); err != nil {
		return nil, kerrors.NewBizStatusError(400, "status is invalid")
	}

	categories, err := String2Category(&req.Categories, &s.ctx)

	if err != nil {
		return nil, err
	}

	p := &model.Product{
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

	if err := model.CreateProduct(mysql.DB, redis.RedisClient, s.ctx, p); err != nil { // 创建商品
		return nil, err
	}
	return &product.CreateProductResp{
		ProductId: p.ID,
		Success:   true,
	}, err
}
