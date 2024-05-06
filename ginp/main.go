package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/file", fileServer)
	r.Run(":9090")
}

func fileServer(c *gin.Context) {
	path := "D:\\work\\code\\go\\tests\\v2"
	fileName := path + c.Query("name")
	c.File(fileName) // 直接c.file 就可以把文件加返回
}
