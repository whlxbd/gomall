package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type PaymentRecord struct {
	gorm.Model
	TransactionId string    `json:"transaction_id"`
	Amount        float32   `json:"amount"`
	OrderId       string    `json:"order_id"`
	UserId        uint32    `json:"user_id"`
	PayAt         time.Time `json:"pay_at"`
}

func (PaymentRecord) TableName() string {
	return "payment_record"
}

func Create(db *gorm.DB, ctx context.Context, pr *PaymentRecord) error {
	return db.WithContext(ctx).Create(pr).Error
}
