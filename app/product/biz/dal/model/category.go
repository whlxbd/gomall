package model

// import (
// 	"context"
// 	"time"
// 	"sync"
// )

type Category struct {
	Base
	Name string `json:"name"`								// 分类名称
	Description string `json:"description"`					// 分类描述

	Products []Product `gorm:"many2many:product_category;"`	// 分类下的商品
}

func (c Category) TableName () string {
	return "category"
}

