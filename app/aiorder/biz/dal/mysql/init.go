package mysql

import (
	"github.com/whlxbd/gomall/app/aiorder/conf"
	"github.com/whlxbd/gomall/app/aiorder/biz/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"fmt"
	"os"
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
	
	if(os.Getenv("GO_ENV") != "online") {
		DB.AutoMigrate(model.Message{})
	}
}
