package casbin_

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/peifengll/tests/testgasbin/self_demo/db"
	"log"
)

var Enforcer *casbin.Enforcer

func InitCasbin() {
	a, err := gormadapter.NewAdapterByDB(db.GetOnceDb())
	if err != nil {
		log.Fatalf("gorm适配器创建失败：%v", err)
	}
	//dir, _ := os.Getwd()
	//modelPath := dir + "/config/rbac_model.conf"
	modelpath := "D:\\code\\go\\trys\\tests\\testgasbin\\self_demo\\config\\rbac_model.conf"
	Enforcer, err = casbin.NewEnforcer(modelpath, a)
	if err != nil {
		log.Fatalf("SetupCasbinEnforcer err: %v", err)
		return
	}
	//Enforcer.LoadPolicy()
	//Enforcer.EnableLog(true)
}
