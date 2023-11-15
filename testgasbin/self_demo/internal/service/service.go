package service

import "github.com/gin-gonic/gin"

func ReadBook(c *gin.Context) {
	c.JSON(200, "成功观看了书")
	return
}

func CloseBook(c *gin.Context) {
	c.JSON(200, "成功关上了书")
	return
}

func OpenBox(c *gin.Context) {
	c.JSON(200, "能打开盒子")
	return
}

func CloseBox(c *gin.Context) {
	c.JSON(200, "能关闭盒子")
	return
}

func ChangeName(c *gin.Context) {
	c.JSON(200, "能更改名字")
	return
}
