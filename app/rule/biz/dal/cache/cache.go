package cache

import (
	"context"
	"sync"

	"github.com/whlxbd/gomall/app/rule/biz/dal/model"
	"github.com/whlxbd/gomall/app/rule/biz/dal/mysql"
	"github.com/whlxbd/gomall/app/rule/biz/dal/redis"
)

var RuleInitOnce sync.Once

func Init() {
	RuleInitOnce.Do(func() {
		// init rule cache
		allRules, err := model.GetAll(mysql.DB, context.Background())
		if err != nil {
			panic(err)
		}
		for _, rule := range allRules {
			ruleKey := "casbin_rule_" + rule.Role
			ok := redis.RedisClient.SAdd(context.Background(), ruleKey, rule.Router)
			if ok.Err() != nil {
				panic(ok)
			}
		}
	})
}

func GetRuleByRole(role string) (*[]string, error) {
	ruleKey := "casbin_rule_" + role
	routers, err := redis.RedisClient.SMembers(context.Background(), ruleKey).Result()
	return &routers, err
}
