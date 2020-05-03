package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 通用结果
func CommonResult(code int, message string, data interface{}) gin.H {
	return gin.H{
		"code":    code,
		"message": message,
		"data":    data,
		"time":    time.Now().Format("2006-01-02 15:04:05:1234"),
	}
}

// 带context 通用结果
func CommonResultWithContext(code int, message string, data interface{}, ctx *gin.Context) {
	ctx.JSON(code, CommonResult(code, message, data))
}

// 错误结果
func ErrorResult(message string) gin.H {
	return CommonResult(http.StatusInternalServerError, message, nil)
}

// 带context通用结果
func ErrorResultWithContext(message string, ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, ErrorResult(message))
}

// 成功结果
func SuccessResult(data interface{}) gin.H {
	return CommonResult(http.StatusOK, "", data)
}

// 带context成功结果
func SuccessResultWithContext(data interface{}, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, SuccessResult(data))
}
