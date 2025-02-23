package whitelist

import (
	"context"

	"gorm.io/gorm"
)

type WhiteRouter struct {
	gorm.Model
	Router string `gorm:"not null"`
}

func (r *WhiteRouter) TableName() string {
	return "white_routers"
}

func Create(db *gorm.DB, ctx context.Context, rule *WhiteRouter) error {
	return db.WithContext(ctx).Create(rule).Error
}

func GetByID(db *gorm.DB, ctx context.Context, id int32) (*WhiteRouter, error) {
	var rule WhiteRouter
	err := db.WithContext(ctx).First(&rule, id).Error
	return &rule, err
}

func GetByRouter(db *gorm.DB, ctx context.Context, router string) (*WhiteRouter, error) {
	var rule WhiteRouter
	err := db.WithContext(ctx).Where("router = ?", router).First(&rule).Error
	return &rule, err
}

func Delete(db *gorm.DB, ctx context.Context, id int32) error {
	return db.WithContext(ctx).Delete(&WhiteRouter{}, id).Error
}

func GetPage(db *gorm.DB, ctx context.Context, page, pageSize int32) ([]*WhiteRouter, error) {
	var rules []*WhiteRouter
	err := db.WithContext(ctx).Offset(int((page - 1) * pageSize)).Limit(int(pageSize)).Find(&rules).Error
	return rules, err
}

func GetByRouterList(db *gorm.DB, ctx context.Context, routers []string) ([]*WhiteRouter, error) {
	var rules []*WhiteRouter
	err := db.WithContext(ctx).Where("router in (?)", routers).Find(&rules).Error
	return rules, err
}
