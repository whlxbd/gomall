package model

import "time"

type Base struct {
	ID uint32 `gorm:"primarykey"`	// 商品ID
	CreatedAt time.Time				// 创建时间
	UpdatedAt time.Time				// 更新时间
}