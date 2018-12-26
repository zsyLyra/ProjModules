package service

import (
	"ProjModules/utils/e"
	"ProjModules/utils/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BindAndValid(ctx *gin.Context, form interface{}) (httpCode int, errCode int) {
	err := ctx.Bind(form)
	if err != nil {
		return http.StatusBadRequest, e.INVALID_PARAMS
	}
	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, e.ERROR
	}
	if !check {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, e.INVALID_PARAMS
	}
	return http.StatusOK, e.SUCCESS
}
