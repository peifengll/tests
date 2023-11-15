package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/peifengll/tests/testgasbin/self_demo/internal/casbin_"
)

func PermissionMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求的path
		p := c.Request.URL.Path
		// 请求的方法
		m := c.Request.Method
		name := c.Request.Header.Get("name")
		fmt.Println("+++++++++++++++++++++++++++++")
		fmt.Printf("path : '%s' , method: '%s' ,name : '%s'\n", p, m, name)
		fmt.Println("+++++++++++++++++++++++++++++")
		isPass, err := casbin_.Enforcer.Enforce(name, p, m)
		if err != nil {
			c.JSON(110, "报错了")
			c.Abort()
			return
		}
		if isPass {
			c.Next()
		} else {
			c.JSON(666, "恭喜，不让访问")
			c.Abort()
			return
		}
	}
}
