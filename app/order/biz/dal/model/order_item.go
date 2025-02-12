package model

import(
	"context"

	"gorm.io/gorm"
)

type OrderItem struct {
	Base
	ProductId    uint32
	OrderIdRefer string `gorm:"size:256;index:,type:btree"` // 添加btree索引类型
	Quantity     int32
	Cost         float32
}

func (oi OrderItem) TableName() string {
	return "order_item"
}

func CreateOrderItems(db *gorm.DB, ctx context.Context, items *[]OrderItem) error {
	return db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(items).Error
	})
}