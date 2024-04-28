package main

import (
	"context"
	"fmt"
	"log"
	"testing"
)

var NotExistsKey = "NotExistsKey"

func TestExistsValue(t *testing.T) {
	r := NewRedis()
	ctx := context.Background()
	result, err := r.Exists(ctx, "666").Result()
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println(result)

}

func TestHMGet(t *testing.T) {
	r := NewRedis()
	ctx := context.Background()
	r.HMSet(ctx, "666", "name", "john", "age", 30)
	result, err := r.HMGet(ctx, "666", "name", "age").Result()
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println(result)
}

func TestMget(t *testing.T) {
	r := NewRedis()
	ctx := context.Background()
	r.Set(ctx, "666", "600", 0)
	r.Set(ctx, "766", "700", 0)
	r.Set(ctx, "866", "800", 0)
	val := []string{"666", "766", "866"}
	result, err := r.MGet(ctx, val...).Result()
	if err != nil {
		log.Fatalln("报错：", err)
	}
	fmt.Println(result)
}
