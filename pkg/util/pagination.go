package util

import (
	"gin-blog/pkg/setting"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// GetPage 分页,保证了各接口的page处理是一致的
func GetPage(c *gin.Context) int {
	res := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		res = (page - 1) * setting.PageSize
	}
	return res
}
