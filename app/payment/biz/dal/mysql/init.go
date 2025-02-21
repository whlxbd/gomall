package mysql

import (
	"fmt"
	"os"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/whlxbd/gomall/app/payment/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/whlxbd/gomall/app/payment/biz/dal/model"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
	klog.Infof("dsn: %s", dsn)
	DB, err = gorm.Open(mysql.Open(dsn),
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
