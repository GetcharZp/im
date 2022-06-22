package middlewares

import (
	"github.com/gin-gonic/gin"
	"im/helper"
	"net/http"
)

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		userClaims, err := helper.AnalyseToken(token)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户认证不通过",
			})
			return
		}
		c.Set("user_claims", userClaims)
		c.Next()
	}
}
