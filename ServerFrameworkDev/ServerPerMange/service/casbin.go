package service

import (
	"ServerFramework/global"
	"strconv"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

type CasbinService struct {
}

var CasbinServiceApp = new(CasbinService)

var CasbinEnforcer *casbin.Enforcer

// 加载casbin
func (receiver *CasbinService) LoadCasbin() *casbin.Enforcer {
	adapter, err := gormadapter.NewAdapterByDB(global.DB)
	if err != nil {
		return nil
	}
	text := `
    [request_definition]
    r = sub, obj, act

    [policy_definition]
    p = sub, obj, act

    [role_definition]
    g = _, _

    [policy_effect]
    e = some(where (p.eft == allow))

    [matchers]
    m = g(r.sub,p.sub) && keyMatch2(r.obj,p.obj) && r.act == p.act
    `
	m, err := model.NewModelFromString(text)
	if err != nil {
		return nil
	}
	CasbinEnforcer, err = casbin.NewEnforcer(m, adapter)
	if err != nil {
		return nil
	}
	CasbinEnforcer.LoadPolicy()
	return CasbinEnforcer
}

func (receiver *CasbinService) AddPolicy(sub, obj, act string) error {
	var policy = gormadapter.CasbinRule{
		Ptype: "p",
		V0:    sub,
		V1:    obj,
		V2:    act,
	}
	return global.DB.Create(&policy).Error
}
func (receiver *CasbinService) RemovePolicy(sub, obj, act string) error {
	return global.DB.Where("v0 = ? AND v1 = ? AND v2 = ?", sub, obj, act).Delete(&gormadapter.CasbinRule{}).Error
}

func (receiver *CasbinService) PolicyExtend(parent, child string) error {
	return global.DB.Create(&gormadapter.CasbinRule{
		Ptype: "g",
		V0:    child,
		V1:    parent,
	}).Error
}

// 
func (receiver *CasbinService) AddRolePolicy(roleID uint64, policies [][]string) error {
	RoldID := strconv.Itoa(int(roleID))
	receiver.ClearCasbin(0, RoldID)
	for _, policy := range policies {
		err := global.DB.Create(&gormadapter.CasbinRule{
			Ptype: "p",
			V0:    RoldID,
			V1:    policy[0],
			V2:    policy[1],
		}).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (receiver *CasbinService) ClearCasbin(v int, p ...string) bool {
	e := receiver.LoadCasbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}
