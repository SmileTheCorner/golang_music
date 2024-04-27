package api

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"go_study/src/common"
	"go_study/src/sys/code_generation/entity"
	"go_study/src/sys/code_generation/service"
	"go_study/src/utils"
	"time"
)

type CodeGenApi struct {
	codeGenService *service.CodeGenService
}

func NewCodeGenApi() *CodeGenApi {
	return &CodeGenApi{
		codeGenService: service.NewCodeGenService(),
	}
}

// 获取当前数据库所有表信息
func (c CodeGenApi) GetCurrentAllTables(ctx *gin.Context) {
	var page common.Page
	err := ctx.ShouldBindJSON(&page)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	tables, count, err := NewCodeGenApi().codeGenService.GetCurrentAllTables(page)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	mapData := utils.PageData(tables, int64(page.PageNum), int64(page.PageSize), count)
	utils.Ok(ctx, mapData)
}

// 导入代码生成业务表
func (c CodeGenApi) ImportGenTable(ctx *gin.Context) {
	var genTables []entity.GenTable
	err := ctx.ShouldBindJSON(&genTables)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	data, err := NewCodeGenApi().codeGenService.ImportGenTable(genTables)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, data)
}

// 分页查询代码生成业务表
func (c CodeGenApi) GetGenTableList(ctx *gin.Context) {
	var page common.Page
	err := ctx.ShouldBindJSON(&page)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	data, count, err := NewCodeGenApi().codeGenService.GetGenTableList(page)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	mapData := utils.PageData(data, int64(page.PageNum), int64(page.PageSize), count)
	utils.Ok(ctx, mapData)
}

// 根据表名查询表信息
func (c CodeGenApi) GetCurrentAllColumns(ctx *gin.Context) {
	tableName := ctx.Query("tableName")
	columns, err := NewCodeGenApi().codeGenService.GetCurrentAllColumns(tableName)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, columns)
}

// 根据ID切片删除代码生成业务表
func (c CodeGenApi) DeleteGenTableByIds(ctx *gin.Context) {
	var ids []string
	err := ctx.ShouldBindJSON(&ids)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	rows, err := NewCodeGenApi().codeGenService.DeleteGenTableByIds(ids)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, rows)
}

// 一键生成代码
func (c CodeGenApi) OneKeyGenerateCode(ctx *gin.Context) {
	var data entity.CodeInfo
	//向data中添加数据
	data.Date = time.Now().Format("2006-01-02 15:04:05")
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	code, err := NewCodeGenApi().codeGenService.GeneratorCode(data)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, code)
}

// 压缩包方式下载
func (c CodeGenApi) DownloadZip(ctx *gin.Context) {
	var ids []string
	err := ctx.ShouldBindJSON(&ids)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	// 创建一个内存缓冲区来存储zip内容
	buf := new(bytes.Buffer)
	err = NewCodeGenApi().codeGenService.GetGenTableByIds(ids, buf)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	//// 设置响应头，告诉浏览器这是一个ZIP文件
	//ctx.Header("Content-Type", "application/zip")
	//ctx.Header("Content-Disposition", "attachment; filename=download.zip")
	//// 将ZIP文件内容写入响应体
	//ctx.Writer.WriteHeader(http.StatusOK)
	//ctx.Writer.Write(buf.Bytes())
	utils.Ok(ctx, buf.Bytes())
}
