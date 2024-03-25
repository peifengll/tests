package main

import "fmt"

type name struct {
}

func main() {
	var i interface{}
	// 检查 interface 变量是否为 nil
	if i == nil {
		fmt.Println("i 是 nil")
	} else {
		fmt.Println("i 不是 nil")
	}
	i = nil
	if i == nil {
		fmt.Println("i 是 nil")
	} else {
		fmt.Println("i 不是 nil")
	}
}
