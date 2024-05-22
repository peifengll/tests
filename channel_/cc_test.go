package channel_

import (
	"fmt"
	"testing"
)

func TestSolve1_1(t *testing.T) {
	solve1_1()
}

func TestSolve1_2(t *testing.T) {
	solve1_2()
}

func TestSolve2_1(t *testing.T) {
	solve2_1()
}

// 测试 长度为n的管道的大小
// 结论： len输出的是管道队列里边现在有几个待读取
func TestLenChannel(t *testing.T) {
	ch := make(chan int, 7)
	ch <- 2
	ch <- 2
	ch <- 2
	fmt.Println(len(ch))
}

/*
=== RUN   TestLenChannel
3
--- PASS: TestLenChannel (0.00s)
PASS
*/
