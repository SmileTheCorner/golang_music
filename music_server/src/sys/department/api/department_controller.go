package api

import (
	"github.com/gin-gonic/gin"
	"go_study/src/common"
	"go_study/src/sys/department/entity"
	"go_study/src/sys/department/service"
	"go_study/src/utils"
	"strconv"
)

type DepartmentApi struct {
	departmentService *service.DepartmentService
}

func NewDepartmentApi() *DepartmentApi {
	return &DepartmentApi{
		departmentService: service.NewDepartmentService(),
	}
}

// 获取部门列表
func (d DepartmentApi) GetDepartmentList(ctx *gin.Context) {
	var page common.Page
	var department entity.Department
	err := ctx.ShouldBind(&department)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	number := ctx.DefaultQuery("pageNum", "1")
	size := ctx.DefaultQuery("pageSize", "10")
	pageNum, _ := strconv.Atoi(number)
	page.PageNum = pageNum
	pageSize, _ := strconv.Atoi(size)
	page.PageSize = pageSize

	lists, count, err := NewDepartmentApi().departmentService.GetDepartmentLists(department, page)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	mapData := utils.PageData(lists, int64(page.PageNum), int64(page.PageSize), count)
	utils.Ok(ctx, mapData)
}

// 添加部门信息
func (d DepartmentApi) AddDepartment(ctx *gin.Context) {
	var department entity.Department
	err := ctx.ShouldBind(&department)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	err = NewDepartmentApi().departmentService.AddDepartment(department)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, nil)
}

// 修改部门信息
func (d DepartmentApi) UpdateDepartment(ctx *gin.Context) {
	var department entity.Department
	err := ctx.ShouldBind(&department)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	err = NewDepartmentApi().departmentService.UpdateDepartment(department)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, nil)
}

// 删除部门信息
func (d DepartmentApi) DropDepartment(ctx *gin.Context) {
	var ids []string
	err := ctx.ShouldBind(&ids)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	err = NewDepartmentApi().departmentService.DropDepartment(ids)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, nil)
}

// 获取部门树形结构
func (d DepartmentApi) GetDepartmentTree(ctx *gin.Context) {
	lists, err := NewDepartmentApi().departmentService.GetDepartmentTree()
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, lists)
}
