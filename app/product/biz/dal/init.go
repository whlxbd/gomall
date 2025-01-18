package dal

import (
	"github.com/whlxbd/gomall/app/product/biz/dal/mysql"
	"github.com/whlxbd/gomall/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
