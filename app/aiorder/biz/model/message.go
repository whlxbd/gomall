package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Message struct {
	Base
	UserId  uint32  `json:"user_id"`
	Content string  `json:"content"`
	OperationType string `json:"operation_type"`
}

func (m *Message) TableName() string {
	return "message"
}

func CreateMessage(ctx context.Context, db *gorm.DB, userId uint32, content string, operationType string) error {
	message := &Message{
		Base:    Base{CreatedAt: time.Now(), UpdatedAt: time.Now()},
		UserId:  userId,
		Content: content,
		OperationType: operationType,
	}

	return db.WithContext(ctx).Create(message).Error
}