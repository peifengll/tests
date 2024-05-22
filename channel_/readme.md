# 问题1 四个协程，每个协程都有自己的编号 ，要求按顺序打印编号并且一直循环

```go

func solve1_1() {
	// id 是编号
	goprint := func(out <-chan struct{}, in chan<- struct{}, id int) {
		nu := struct {
		}{}
		for {
			<-out
			fmt.Println(id)
			in <- nu
		}
	}
	// 初始化四个协程，每个都要等前边的就可以了
	a, b, c, d := make(chan struct{}), make(chan struct{}), make(chan struct{}), make(chan struct{})
	// 四个协程，每个收到了上一个发送的信息 ，就可以了
	go goprint(a, b, 1)
	go goprint(b, c, 2)
	go goprint(c, d, 3)
	go goprint(d, a, 4)
	a <- struct{}{}
	select {}
}

func solve1_2() {
	ch := make([]chan struct{}, 4) // 初始化四个
	for i := 0; i < 4; i++ {
		ch[i] = make(chan struct{})
	}
	for i := 0; i <= 3; i++ {
		go func(k int) {
			for {
				<-ch[k-1]
				fmt.Println(k)
				ch[k%4] <- struct{}{}
			}
		}(i + 1)
	}
	ch[0] <- struct{}{}
	select {}
}

```



# 问题2 用channel 实现一个限流器


```go

```
