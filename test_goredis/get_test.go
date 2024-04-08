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
