package main

import (
	"context"
	"log"
	"testing"
)

func TestListPush(t *testing.T) {
	r := NewRedis()
	defer r.Close()
	ctx := context.Background()
	r.RPush(ctx, "test_ddd", 99)
}

// 获取 列表
func TestLrange(t *testing.T) {
	r := NewRedis()
	defer r.Close()
	ctx := context.Background()
	result, err := r.LRange(ctx, "test_ddd", 0, -1).Result()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(result)
	//log.Println(result[0])
}

func TestLLen(t *testing.T) {
	r := NewRedis()
	defer r.Close()
	ctx := context.Background()
	result, err := r.LLen(ctx, "test_ddd").Result()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(result)
	//log.Println(result[0])
}

// 会取出来一个，result就算取的那一个
func TestLPpo(t *testing.T) {
	r := NewRedis()
	defer r.Close()
	ctx := context.Background()
	result, err := r.LPop(ctx, "test_ddd").Result()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(result)
	//log.Println(result[0])
}

// 会取出来 min(count,len)个，result就算取的那几个
func TestLPopcount(t *testing.T) {
	r := NewRedis()
	defer r.Close()
	ctx := context.Background()
	result, err := r.LPopCount(ctx, "test_ddd", 1).Result()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(result)
	//log.Println(result[0])
}
