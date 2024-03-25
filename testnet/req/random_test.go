package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRandom(t *testing.T) {
	// 设置种子
	rand.Seed(time.Now().UnixNano())

	// 生成随机数
	randomNum := rand.Intn(30) // 生成 0 到 100 之间的随机整数
	fmt.Println("随机数:", randomNum)
}
