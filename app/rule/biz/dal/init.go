package dal

import (
	"github.com/whlxbd/gomall/app/rule/biz/dal/mysql"
	"github.com/whlxbd/gomall/app/rule/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
