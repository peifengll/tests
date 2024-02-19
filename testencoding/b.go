package main

import (
	"bytes"
	"fmt"
)

//bytes.Buffer 进行多个bytes得连接

func main() {
	b1 := []byte("this is a first string")

	b2 := []byte("this is a second string")
	var buffer bytes.Buffer
	buffer.Write(b1)
	buffer.Write(b2)
	b3 := buffer.Bytes()
	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(string(b3))
}
