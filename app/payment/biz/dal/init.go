package dal

import (
	"github.com/whlxbd/gomall/app/payment/biz/dal/mysql"
	"github.com/whlxbd/gomall/app/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
