package model

import (
	"context"
	"gorm.io/gorm"
)

type User struct {
	Base
	Email    string `gorm:"unique;not null" validate:"required,email"`
	Password string `gorm:"not null"`
}

func (u *User) TableName() string {
	return "users"
}

func Create(db *gorm.DB, ctx context.Context, user *User) error {
	return db.WithContext(ctx).Create(user).Error
}

func GetByEmail(db *gorm.DB, ctx context.Context, email string) (*User, error) {
	var user User
	err := db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	return &user, err
}
