package dal

import (
	"github.com/whlxbd/gomall/rpc_gen/biz/dal/mysql"
	"github.com/whlxbd/gomall/rpc_gen/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
