package model

import (
	"context"

	"gorm.io/gorm"
)

type Rule struct {
	gorm.Model
	Role   string `gorm:"not null"`
	Router string `gorm:"not null"`
}

func (r *Rule) TableName() string {
	return "rules"
}

func Create(db *gorm.DB, ctx context.Context, rule *Rule) error {
	return db.WithContext(ctx).Create(rule).Error
}

func GetByID(db *gorm.DB, ctx context.Context, id int32) (*Rule, error) {
	var rule Rule
	err := db.WithContext(ctx).First(&rule, id).Error
	return &rule, err
}

func GetByRole(db *gorm.DB, ctx context.Context, role string) (*Rule, error) {
	var rule Rule
	err := db.WithContext(ctx).Where("role = ?", role).First(&rule).Error
	return &rule, err
}

func GetByRouter(db *gorm.DB, ctx context.Context, router string) (*Rule, error) {
	var rule Rule
	err := db.WithContext(ctx).Where("router = ?", router).First(&rule).Error
	return &rule, err
}

func Update(db *gorm.DB, ctx context.Context, rule *Rule) error {
	return db.WithContext(ctx).Model(rule).Updates(rule).Error
}

func Delete(db *gorm.DB, ctx context.Context, id int32) error {
	return db.WithContext(ctx).Delete(&Rule{}, id).Error
}

func GetPage(db *gorm.DB, ctx context.Context, page, pageSize int32) ([]*Rule, error) {
	var rules []*Rule
	err := db.WithContext(ctx).Limit(int(pageSize)).Offset(int(page * pageSize)).Find(&rules).Error
	return rules, err
}

func GetAll(db *gorm.DB, ctx context.Context) ([]*Rule, error) {
	var rules []*Rule
	err := db.WithContext(ctx).Find(&rules).Error
	return rules, err
}
