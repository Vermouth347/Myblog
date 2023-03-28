package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 在前后端交互时，我们还需解决跨域问题
// CORSMiddleware是一个跨域请求的中间件，来处理跨域请求
/* middleware/CORSMiddleware.go/ */
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
