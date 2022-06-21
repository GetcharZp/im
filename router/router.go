package router

import (
	"github.com/gin-gonic/gin"
	"im/service"
)

func Router() *gin.Engine {
	r := gin.Default()

	// 用户登录
	r.POST("/login", service.Login)

	return r
}
