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
	// 发送验证码
	r.POST("/send/code", service.SendCode)

	auth := r.Group("/u", middlewares.AuthCheck())

	// 用户详情
	auth.GET("/user/detail", service.UserDetail)
	// 发送、接受消息
	auth.GET("/websocket/message", service.WebsocketMessage)

	return r
}
