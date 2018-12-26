package service

import (
	"ProjModules/utils/e"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	Ctx *gin.Context
}

func (g *Gin) Response(httpCode int, errCode int, data interface{})  {
	g.Ctx.JSON(httpCode, gin.H{
		"code" : errCode,
		"msg" : e.GetMsg(errCode),
		"data" : data,
	})
	return
}
