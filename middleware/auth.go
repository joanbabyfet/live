package middleware

import (
	"github.com/gin-gonic/gin"
)

//模拟登录
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetHeader("X-User-ID")

		if userID == "" {
			c.JSON(401, gin.H{
				"error": "unauthorized",
			})

			c.Abort()
			return
		}
		
		//存入上下文, 上下文（Context）就是一次 HTTP 请求的“共享数据容器”
		c.Set("userID", userID)

		c.Next()
	}
}