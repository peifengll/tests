package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type RedisObj struct {
	*redis.Client
}

func NewRedis() RedisObj {
	client := redis.NewClient(&redis.Options{
		Addr:     "peifeng.site:6379", // redis地址
		Password: "ningzaichun",       // 密码
		DB:       0,                   // 使用默认数据库
	})
	return RedisObj{client}
}

type AutoGenerated struct {
	DpvMatchCode string `json:"dpv_match_code"`
	DpvFootnotes string `json:"dpv_footnotes"`
	DpvCmra      string `json:"dpv_cmra"`
	DpvVacant    string `json:"dpv_vacant"`
	Active       string `json:"active"`
}

func (m *AutoGenerated) MarshalBinary() (data []byte, err error) {
	return json.Marshal(m)
}

func (m *AutoGenerated) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)

}

func main() {
	r := NewRedis()
	ctx := context.Background()
	a := &AutoGenerated{
		DpvMatchCode: "Y",
		DpvFootnotes: "AABB",
		DpvCmra:      "N",
		DpvVacant:    "N",
		Active:       "Y",
	}
	len, err := r.RPush(ctx, "666", a).Result()
	fmt.Println(len)
	fmt.Println(err)

	//var b AutoGenerated
	//err := r.RPop(ctx, "666").Scan(&b)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(b)
}
