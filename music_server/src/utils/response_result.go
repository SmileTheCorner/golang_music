package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Ok 操作成功
func Ok(ctx *gin.Context, data interface{}) {
	ctx.AbortWithStatusJSON(http.StatusOK, &ResponseResult{
		Code: 200,
		Msg:  "操作成功",
		Data: data,
	})
}

// Fail 失败操作
func Fail(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, &ResponseResult{
		Code: 400,
		Msg:  err.Error(),
		Data: nil,
	})
}

func ResponseCustom(ctx *gin.Context, code int, msg string, data interface{}) {
	ctx.AbortWithStatusJSON(code, &ResponseResult{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
