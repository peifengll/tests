package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/peifengll/tests/testgasbin/self_demo/internal/middleware"
	"github.com/peifengll/tests/testgasbin/self_demo/internal/service"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.PermissionMiddleWare())
	//	book box name
	router.GET("/readbook", service.ReadBook)
	router.POST("/closebook", service.CloseBook)
	router.POST("/openbox", service.OpenBox)
	router.POST("/closebox", service.CloseBox)
	router.POST("/changename", service.ChangeName)
	return router
}
