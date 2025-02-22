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
	"github.com/whlxbd/gomall/app/checkout/biz/dal/mysql"
	"gorm.io/gorm"

	rulemodel "github.com/whlxbd/gomall/app/rule/biz/dal/model"
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

		err = casbinEnf.LoadPolicy()
		if err != nil {
			panic(err)
		}
	})
}

func ReadPolicyFromDB() error {
	Init()
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
	return mysql.DB.Transaction(func(tx *gorm.DB) error {
		rule := &rulemodel.Rule{
			Role:   sub,
			Router: act,
		}
		err := rulemodel.Create(tx, context.Background(), rule)
		if err != nil {
			return err
		}

		ok, err := casbinEnf.AddPolicy(sub, act)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("add policy failed")
		}
		return nil
	})
}

func RemovePolicy(sub string, act string) error {
	Init()
	return mysql.DB.Transaction(func(tx *gorm.DB) error {
		rule, err := rulemodel.GetByRoleAndRouter(mysql.DB, context.Background(), sub, act)
		if err != nil {
			return err
		}

		if rule.Router != act {
			return errors.New("rule not found")
		}

		err = rulemodel.Delete(tx, context.Background(), int32(rule.ID))
		if err != nil {
			return err
		}

		ok, err := casbinEnf.RemovePolicy(sub, act)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("remove policy failed")
		}
		return nil
	})
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
