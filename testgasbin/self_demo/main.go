package main

import (
	"github.com/peifengll/tests/testgasbin/self_demo/config"
	v1 "github.com/peifengll/tests/testgasbin/self_demo/internal/apiserver/v1"
	"github.com/peifengll/tests/testgasbin/self_demo/internal/casbin_"
)

func main() {
	config.InitConfig()
	casbin_.InitCasbin()
	r := v1.Router()
	r.Run(":8080")
}

//http://127.0.0.1:8080/readbook
//http://127.0.0.1:8080/closebook
//http://127.0.0.1:8080/openboox
//http://127.0.0.1:8080/closebox
//http://127.0.0.1:8080/changename
