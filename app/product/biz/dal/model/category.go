package model

type Category struct {
	Base
	Name string `json:"name"`
	Description string `json:"description"`

	Products []Product `gorm:"many2many:product_category;"`
}

func (c Category) TableName () string {
	return "category"
}
