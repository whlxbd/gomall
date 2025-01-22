package model

import (
	"time"

	"gorm.io/gorm"
)

type PaymentRecord struct {
	gorm.Model
	Amount float32
	OrderId string
	UserId uint32
	PayAt time.Time
}