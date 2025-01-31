package model

import (
	
)

type Cart struct {
	Base
	UserID uint32 `gorm:"user_id"`	// 用户ID
	ProductID uint32 `gorm:"product_id"`	// 商品ID
	Quantity uint32 `gorm:"quantity"`	// 商品数量
}

func (c Cart) TableName() string {
	return "cart"
}

