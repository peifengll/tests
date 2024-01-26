package router

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/peifengll/tests/UEditor/controller"
)

func Server() *gin.Engine {
	server := gin.Default()
	server.Static("/static/", "./static")
	g := server.Group("/ueditor")
	g.Any("/hello", controllers.Action)
	return server
}
