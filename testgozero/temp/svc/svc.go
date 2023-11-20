package svc

import (
	"fmt"
	"github.com/peifengll/tests/testgozero/temp/config"
	"github.com/peifengll/tests/testgozero/temp/model/have_c"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var RedisClient *redis.Redis
var ScoreModel have_c.ScoreModel

func Init() {
	c, err := config.GetConfig()
	if err != nil {
		fmt.Println("配置不正常")
		return
	}
	RedisClient = redis.New(c.Redis.Host, func(r *redis.Redis) {
		r.Type = c.Redis.Type
		r.Pass = c.Redis.Pass
	})
	mysqlConn := sqlx.NewMysql("root:peifeng@(127.0.0.1:3306)/test_1?charset=utf8&parseTime=true")
	have_c.NewScoreModel(mysqlConn, c.CacheRedis)
}
