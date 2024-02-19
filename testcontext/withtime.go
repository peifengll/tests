package main

import (
	"context"
	"time"
)

func solve() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	for i := 0; i <= 15; i++ {
		select {
		case <-ctx.Done():
			println("timeout")
			return
		default:
			time.Sleep(1 * time.Second)
			println(i)
		}
	}
	println("done")
}

func solve1() {
	ch := make(chan struct{})
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(1 * time.Second)
			println(i)
		}
		ch <- struct{}{}
	}()
	select {
	case <-ch:
		println("正常结束")
	case <-time.After(10 * time.Second):
		println("超时了")
	}
}
func main() {
	//solve()
	solve1()
}
