package model

import (
	product "github.com/whlxbd/gomall/app/product/biz/dal/model"
)

type Cart struct {
    Base
    UserID    uint32    `json:"user_id" gorm:"index:idx_user_product;not null;comment:用户ID"`
    ProductID uint32    `json:"product_id" gorm:"index:idx_user_product;not null;comment:商品ID"`
    Quantity  uint32    `json:"quantity" gorm:"not null;default:1;check:quantity > 0;comment:商品数量"`
    Selected  bool      `json:"selected" gorm:"default:true;comment:是否选中"`
    Status    int8      `json:"status" gorm:"default:1;comment:状态 1:正常 2:失效"`
    
    // 关联商品信息(预加载)
    Product   *product.Product  `json:"product" gorm:"foreignKey:ProductID"`
}

func (c Cart) TableName() string {
	return "cart"
}

