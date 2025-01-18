package dal

import (
	"github.com/whlxbd/gomall/app/cart/biz/dal/mysql"
	"github.com/whlxbd/gomall/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
