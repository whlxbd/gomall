package model

import (
	"context"

	"gorm.io/gorm"
)

type Consignee struct {
	Email string

	StreetAddress string
	City          string
	State         string
	Country       string
	ZipCode       int32
}

type OrderState string

const (
	OrderStatePlaced   OrderState = "placed"
	OrderStatePaid     OrderState = "paid"
	OrderStateCanceled OrderState = "canceled"
)

type Order struct {
    Base
    OrderId      string     `gorm:"uniqueIndex;size:256"` 
    UserId       uint32     `gorm:"index"`
    UserCurrency string
    Consignee    Consignee  `gorm:"embedded"`
    OrderItems   []OrderItem `gorm:"foreignKey:OrderIdRefer;references:OrderId"`
    OrderState   OrderState `gorm:"index"`
}

func (o Order) TableName() string {
	return "order"
}

func ListOrder(db *gorm.DB, ctx context.Context, userId uint32) (orders []Order, err error) {
	err = db.Model(&Order{}).Where(&Order{UserId: userId}).Preload("OrderItems").Find(&orders).Error
	return
}

func GetOrder(db *gorm.DB, ctx context.Context, userId uint32, orderId string) (order Order, err error) {
	err = db.Where(&Order{UserId: userId, OrderId: orderId}).First(&order).Error
	return
}

func UpdateOrderState(db *gorm.DB, ctx context.Context, userId uint32, orderId string, state OrderState) error {
	return db.Transaction(func(tx *gorm.DB) error {
		order := Order{}
		if err := tx.Where(&Order{UserId: userId, OrderId: orderId}).First(&order).Error; err != nil {
			return err
		}
		order.OrderState = state
		return tx.Save(&order).Error
	})
}

func CreateOrder(db *gorm.DB, ctx context.Context, order *Order) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(order).Error; err != nil {
			return err
		}
		return nil
	})
}

func EditOrder(db *gorm.DB, ctx context.Context, order *Order) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(order).Error; err != nil {
			return err
		}
		return nil
	})
}