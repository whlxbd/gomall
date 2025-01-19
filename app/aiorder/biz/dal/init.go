package dal

import (
	"github.com/whlxbd/gomall/app/aiorder/biz/dal/mysql"
	"github.com/whlxbd/gomall/app/aiorder/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
