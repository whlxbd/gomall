package cas

import (
	"context"
	"errors"
	"os"
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	redisadapter "github.com/casbin/redis-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"github.com/whlxbd/gomall/app/rule/biz/dal/mysql"

	rulemodel "github.com/whlxbd/gomall/app/rule/biz/dal/model/rule"
)

var (
	casbinInit sync.Once
	casbinEnf  *casbin.Enforcer
)

func Init() {
	casbinInit.Do(func() {
		a, err := redisadapter.NewAdapter("tcp", os.Getenv("REDIS_ADDR"))
		if err != nil {
			panic(err)
		}
		m, err := model.NewModelFromString(`[request_definition]
			r = sub, act

			[policy_definition]
			p = sub, act

			[policy_effect]
			e = some(where (p.eft == allow))

			[matchers]
			m = r.sub == p.sub && r.act == p.act`)
		if err != nil {
			panic(err)
		}

		casbinEnf, err = casbin.NewEnforcer(m, a)
		if err != nil {
			panic(err)
		}

		err = readPolicyFromDB()
		if err != nil {
			panic(err)
		}
	})
}

func readPolicyFromDB() error {
	rules, err := rulemodel.GetAll(mysql.DB, context.Background())
	if err != nil {
		return err
	}

	for _, rule := range rules {
		ok, err := casbinEnf.AddPolicy(rule.Role, rule.Router)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("add policy failed")
		}
	}

	return nil
}

func AddPolicy(sub string, act string) error {
	Init()
	ok, err := casbinEnf.AddPolicy(sub, act)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("add policy failed")
	}
	return nil
}

func RemovePolicy(sub string, act string) error {
	Init()
	ok, err := casbinEnf.RemovePolicy(sub, act)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("remove policy failed")
	}
	return nil
}

func CheckPolicy(sub string, act string) error {
	Init()
	ok, err := casbinEnf.Enforce(sub, act)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("permission denied")
	}
	return nil
}
