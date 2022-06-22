package router

import (
	"github.com/gin-gonic/gin"
	"im/middlewares"
	"im/service"
)

func Router() *gin.Engine {
	r := gin.Default()

	// 用户登录
	r.POST("/login", service.Login)

	auth := r.Group("/u", middlewares.AuthCheck())

	// 用户详情
	auth.GET("/user/detail", service.UserDetail)

	return r
}
