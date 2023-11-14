package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"log"
)

func check(e *casbin.Enforcer, sub, obj, act string) {
	ok, _ := e.Enforce(sub, obj, act)
	if ok {
		fmt.Printf("%s CAN %s %s\n", sub, act, obj)
	} else {
		fmt.Printf("%s CANNOT %s %s\n", sub, act, obj)
	}
}

func main() {
	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		log.Fatalf("NewEnforecer failed:%v\n", err)
	}
	checker1(e)
}
func checker1(e *casbin.Enforcer) {
	check(e, "dajun", "prod.data", "read")
	check(e, "dajun", "prod.data", "write")
	check(e, "lizi", "dev.data", "read")
	check(e, "lizi", "dev.data", "write")
	check(e, "lizi", "prod.data", "write")
}
