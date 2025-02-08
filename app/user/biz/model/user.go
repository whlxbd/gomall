package model

import (
	"context"
	"gorm.io/gorm"
)

type User struct {
	Base
	Email    string `gorm:"unique;not null" validate:"required,email"`
	Password string `gorm:"not null"`
	UserName string `gorm:"not null"`
	AvatarUrl string `gorm:"not null"`
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

func GetByID(db *gorm.DB, ctx context.Context, id int32) (*User, error) {
	var user User
	err := db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	return &user, err
}

func GetByUserName(db *gorm.DB, ctx context.Context, userName string) (*User, error) {
	var user User
	err := db.WithContext(ctx).Where("user_name = ?", userName).First(&user).Error
	return &user, err
}

func Update(db *gorm.DB, ctx context.Context, user *User) error {
	return db.WithContext(ctx).Model(user).Updates(user).Error
}

func Delete(db *gorm.DB, ctx context.Context, id int32) error {
	return db.WithContext(ctx).Delete(&User{}, id).Error
}
