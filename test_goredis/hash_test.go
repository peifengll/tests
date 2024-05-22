package main

import (
	"context"
	"log"
	"strconv"
	"testing"
)

func TestShouldBind(t *testing.T) {
	r := NewRedis()
	ctx := context.Background()
	key := "space_repeat_BLNOXd36vf_rate"
	res, err := r.HGetAll(ctx, key).Result()

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res)

}

func TestLoadParm(t *testing.T) {
	r := NewRedis()
	ctx := context.Background()
	key := "test_dic_parm"
	r.HSet(ctx, key,
		"max_interval", 362,
		//"weights", "0.4, 0.6, 2.4, 5.8, 4.93, 0.94, 0.86, 0.01, 1.49, 0.14, 0.94, 2.18, 0.05, 0.34, 1.26, 0.29, 2.61",
		"request_retention", 0.9)
	dic := r.HGetAll(ctx, key).Val()
	log.Println(dic)
	float1, err := strconv.ParseFloat(dic["max_interval"], 10)
	if err != nil {
		log.Fatalln(err)
	}

	RequestRetention, err := strconv.ParseFloat(dic["request_retention"], 10)
	log.Println(float1, RequestRetention)
}
