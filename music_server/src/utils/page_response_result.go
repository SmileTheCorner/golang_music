package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PageResponseResult struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data"`
}

// PageData 封装分页返回数据
func PageData(list interface{}, pageNum, pageSize, total int64) map[string]interface{} {
	mapData := make(map[string]interface{}, 5)
	mapData["list"] = list
	mapData["pageNum"] = pageNum
	mapData["pageSize"] = pageSize
	mapData["total"] = total
	return mapData
}

// Ok 操作成功
func (m PageResponseResult) Ok(ctx *gin.Context, data interface{}) {
	ctx.AbortWithStatusJSON(http.StatusOK, &ResponseResult{
		Code: 200,
		Msg:  "操作成功",
		Data: data,
	})
}

// Fail 失败操作
func (m PageResponseResult) Fail(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, &ResponseResult{
		Code: 400,
		Msg:  err.Error(),
		Data: nil,
	})
}
