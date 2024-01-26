package main

import (
	"bytes"
	"log"
	"os"
)

var configJson []byte // 当客户端请求 /ueditor/go/controller?action=config 返回的json内容

func init() {
	file, err := os.Open("D:\\code\\go\\trys\\tests\\others\\Go-UEditor\\conf\\config.json")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer file.Close()
	buf := bytes.NewBuffer(nil)
	buf.ReadFrom(file)

	configJson = buf.Bytes()
}
