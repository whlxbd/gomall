package dal

import (
	"github.com/whlxbd/gomall/app/order/biz/dal/mysql"
	"github.com/whlxbd/gomall/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
