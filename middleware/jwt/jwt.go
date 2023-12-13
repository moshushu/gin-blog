package jwt

import (
	"gin-blog/pkg/e"
	"gin-blog/pkg/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// JWT 中间价
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				// 验证token是否过期
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			// 终止请求
			c.Abort()
			return
		}

		// 继续请求
		c.Next()
	}
}
