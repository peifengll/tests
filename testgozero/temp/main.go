package main

import (
	"context"
	"fmt"
	"github.com/peifengll/tests/testgozero/temp/config"
	"github.com/peifengll/tests/testgozero/temp/svc"
)

func main() {
	config.InitConfig()
	svc.Init()
	//c, err := config.GetConfig()
	//if err != nil {
	//	log.Fatal("err is ", err)
	//}
	//fmt.Println(c)
	//fmt.Printf("%+v", c.Cache)
	ctx := context.Background()
	var err error

	// 本地
	//one, err := svc.SqlcScoreModel.FindOne(ctx, 1)
	//one, err = svc.SqlcScoreModel.FindOne(ctx, 3)
	//one, err = svc.SqlcScoreModel.FindOne(ctx, 5)
	//one, err = svc.SqlcScoreModel.FindOne(ctx, 7)
	//阿里
	//one, err := svc.SqlcScoreModel.FindOne(ctx, 4)
	//one, err = svc.SqlcScoreModel.FindOne(ctx, 6)
	//one, err = svc.SqlcScoreModel.FindOne(ctx, 8)

	err = svc.SqlcScoreModel.Delete(ctx, 8)

	if err != nil {
		fmt.Println("err is ", err)
		return
	}
	//fmt.Println("数据是:", one)
}
