package dal

import (
	"github.com/whlxbd/gomall/app/cart/biz/dal/mysql"
	"github.com/whlxbd/gomall/app/cart/biz/dal/redis"
	// "github.com/whlxbd/gomall/app/cart/biz/dal/mq"
)

func Init() {
	redis.Init()
	mysql.Init()
	// err := mq.Init(mysql.DB, redis.RedisClient)
	// if err != nil {
	// 	panic(err)
	// }
}
