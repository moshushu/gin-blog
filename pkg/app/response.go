package app

import (
	"github.com/gin-gonic/gin"
	"github.com/moshushu/gin-blog/pkg/e"
)

type Gin struct {
	C *gin.Context
}

// 这样子以后如果要变动，直接改动 app 包内的方法即可
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": errCode,
		"msg":  e.GetMsg(errCode),
		"data": data,
	})
}
