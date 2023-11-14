package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"log"
)

//ACL（access-control-list，访问控制列表）。ACL显示定义了每个主体对每个资源的权
//限情况， 未定义的就没有权限。我们还可以加上超级管理员，超级管理员可以进行任何操作。假设超级管理员为root，
//我们只需要修改匹配器：

func check(e *casbin.Enforcer, sub, obj, act string) {
	ok, _ := e.Enforce(sub, obj, act)
	if ok {
		fmt.Printf("%s CAN %s %s \n", sub, act, obj)
	} else {
		fmt.Printf("%s CANNOT %s %s \n", sub, act, obj)
	}
}

func main() {
	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		log.Fatalf("NewEnforecer failed:%v \n", err)
	}
	//checker1(e)
	checker2(e)

}

// 测试多个人
func checker1(e *casbin.Enforcer) {
	check(e, "dajun", "data1", "read")
	check(e, "lizi", "data2", "write")
	check(e, "dajun", "data1", "write")
	check(e, "dajun", "data2", "read")
}

// 测试root
func checker2(e *casbin.Enforcer) {
	check(e, "root", "data1", "read")
	check(e, "root", "data2", "write")
	check(e, "root", "data1", "execute")
	check(e, "root", "data2", "rwx")
}
