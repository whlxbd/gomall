package dal

import (
	"github.com/whlxbd/gomall/app/auth/biz/dal/mysql"
	"github.com/whlxbd/gomall/app/auth/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
