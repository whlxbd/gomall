package model

import (
	"context"

	"gorm.io/gorm"
)

type Product struct {
	Base
	Name string `json:"name"`
	Description string `json:"description"`
	Picture string `json:"picture"`
	Price float32 `json:"price"`

	Categories []Category `gorm:"many2many:product_category;"`
}

type ProductQuery struct {
	ctx context.Context
	db *gorm.DB
}

func (p Product) TableName () string {
	return "product"
}


func (p ProductQuery) GetById (productid int32) (product Product, err error) {
	return nil, nil
}