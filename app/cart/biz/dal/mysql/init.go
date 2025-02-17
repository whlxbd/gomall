package mysql

import (
	"fmt"
	"os"

	"github.com/whlxbd/gomall/app/cart/biz/dal/model"
	"github.com/whlxbd/gomall/app/cart/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
		needDemoData := !DB.Migrator().HasTable(&model.Cart{})
		err := DB.AutoMigrate( //nolint:errcheck
			&model.Cart{},
		)
		if needDemoData {
			
		}

		if err != nil {
			panic(err)
		}
	}
}
