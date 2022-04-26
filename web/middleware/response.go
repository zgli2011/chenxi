package middleware

import (
	"chenxi/initialize"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS = 0
	FAILED  = 1
)

func ResponsePageDataSuccess(ctx *gin.Context, data interface{}, count *int64, page_size interface{}, msg string) {
	page_total := int(math.Ceil(float64(int(*count)) / float64(page_size.(int))))
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": SUCCESS,
		"msg":  msg,
		"info": map[string]interface{}{
			"count":         count,
			"page_total":    page_total,
			"page_size":     page_size.(int),
			"max_page_size": initialize.Config.System.PageSize,
			"results":       data,
		},
	})
}

func ResponseSuccess(ctx *gin.Context, data interface{}, msg string) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": SUCCESS,
		"msg":  msg,
		"info": data,
	})
}

func ResponseFailed(ctx *gin.Context, msg string) {
	// if msg == "record not found" {
	// 	http_code = http.StatusNotFound
	// }
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": FAILED,
		"info": nil,
		"msg":  msg,
	})
}
