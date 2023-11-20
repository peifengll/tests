package main

import (
	"context"
	"fmt"
	"github.com/peifengll/tests/testgozero/temp/config"
	"github.com/peifengll/tests/testgozero/temp/svc"
	"log"
)

func main() {
	config.InitConfig()
	svc.Init()
	c, err := config.GetConfig()
	if err != nil {
		log.Fatal("err is ", err)
	}
	fmt.Println(c)
	fmt.Printf("%+v", c.Cache)
	ctx := context.Background()
	one, err := svc.SqlcScoreModel.FindOne(ctx, 0)
	if err != nil {
		return
	}
	fmt.Println(one)
}
