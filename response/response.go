package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//为减少代码量，我们为项目封装一个统一的失败与成功的返回格式

func Response(c *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	c.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}

// Success 成功
func Success(c *gin.Context, data gin.H, msg string) {
	Response(c, http.StatusOK, 200, data, msg)
}

// Fail 失败
func Fail(c *gin.Context, data gin.H, msg string) {
	Response(c, http.StatusOK, 400, data, msg)
}
