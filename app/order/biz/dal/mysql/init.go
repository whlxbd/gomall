package mysql

import (
	"fmt"
	"os"

	"github.com/whlxbd/gomall/app/order/biz/dal/model"
	"github.com/whlxbd/gomall/app/order/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "github.com/brianvoe/gofakeit/v7"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	conf.GetConf().MySQL.DSN = dsn
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	if os.Getenv("GO_ENV") != "online" {
		// 先创建订单表
		if err := DB.AutoMigrate(&model.Order{}); err != nil {
			panic(err)
		}

		// 再创建关联表
		if err := DB.AutoMigrate(
			&model.OrderItem{},
		); err != nil {
			panic(err)
		}
	}
}
