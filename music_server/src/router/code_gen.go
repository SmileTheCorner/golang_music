package router

import (
	"github.com/gin-gonic/gin"
	"go_study/src/sys/code_generation/api"
)

func InitCodeGenRouter() {
	RegistRoute(func(rgPblicRoute *gin.RouterGroup, rgAuthRoute *gin.RouterGroup) {
		genApi := api.NewCodeGenApi()
		group := rgAuthRoute.Group("codeGen")
		{
			//获取所有表信息
			group.POST("/allTables", genApi.GetCurrentAllTables)
			//一键生成代码
			group.POST("/oneKeyGenerateCode", genApi.OneKeyGenerateCode)
			//导入代码生成业务表
			group.POST("/importGenTable", genApi.ImportGenTable)
			//分页查询代码生成业务表
			group.POST("/genTableList", genApi.GetGenTableList)
			//根据表名查询表信息
			group.GET("/allColumnsByTableName", genApi.GetCurrentAllColumns)
			// 根据ID切片删除代码生成业务表
			group.DELETE("/dropGenTableByIds", genApi.DeleteGenTableByIds)
			//下载生成的代码
			group.POST("/downloadCode", genApi.DownloadZip)
		}
	})
}
