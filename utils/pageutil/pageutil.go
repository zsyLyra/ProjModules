package pageutil

import (
	"ProjModules/utils/setting"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
)

func GetPage(ctx *gin.Context) int {
	result := 0
	page, _ := com.StrTo(ctx.Query("page")).Int()
	if page > 0 {
		result = (page-1)*setting.AppSetting.PageSize
	}
	return result
}
