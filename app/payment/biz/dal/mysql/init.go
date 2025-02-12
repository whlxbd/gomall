package mysql

import (
	"github.com/whlxbd/gomall/app/payment/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/whlxbd/gomall/app/payment/biz/model"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&model.PaymentRecord{})
}
