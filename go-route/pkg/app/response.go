package app

import (
	"book-server/pkg/e"

	"github.com/gin-gonic/gin"
)

//Gin Gin
type Gin struct {
	C *gin.Context
}

//Respinse Response
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": httpCode,
		"msg":  e.GetMsg(errCode),
		"data": data,
	})
	return
}
