package dal

import (
	"github.com/whlxbd/gomall/app/checkout/biz/dal/mysql"
	"github.com/whlxbd/gomall/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
