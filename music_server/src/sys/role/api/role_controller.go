package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go_study/src/common"
	"go_study/src/sys/role/entity"
	"go_study/src/sys/role/service"
	"go_study/src/utils"
	"strconv"
)

type RoleApi struct {
	roleService *service.RoleService
}

func NewRoleApi() RoleApi {
	return RoleApi{
		roleService: service.NewRoleService(),
	}
}

// 获取角色列表
func (m RoleApi) GetRoleList(ctx *gin.Context) {
	var page common.Page
	var role entity.Role
	err := ctx.ShouldBind(&role)
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

	lists, count, err := NewRoleApi().roleService.GetRoleLists(role, page)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	var roleList []entity.RoleReq
	//将部门数据序列化成切片
	for _, item := range lists {
		var roleReq entity.RoleReq
		roleReq.ID = item.ID
		roleReq.Name = item.Name
		roleReq.Code = item.Code
		roleReq.Sort = item.Sort
		roleReq.DataScope = item.DataScope
		roleReq.Status = item.Status
		roleReq.Remark = item.Remark
		roleReq.CreateBy = item.CreateBy
		roleReq.UpdateBy = item.UpdateBy
		json.Unmarshal([]byte(item.DataScopeDept), &roleReq.DataScopeDept)
		roleList = append(roleList, roleReq)
	}
	mapData := utils.PageData(roleList, int64(page.PageNum), int64(page.PageSize), count)
	utils.Ok(ctx, mapData)
}

// 添加角色信息
func (m RoleApi) AddRole(ctx *gin.Context) {
	var role entity.RoleReq
	err := ctx.ShouldBind(&role)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	err = NewRoleApi().roleService.AddRole(role)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, nil)
}

// 修改角色信息
func (m RoleApi) UpdateRole(ctx *gin.Context) {
	var role entity.RoleReq
	err := ctx.ShouldBind(&role)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	err = NewRoleApi().roleService.UpdateRole(role)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, nil)
}

// 删除角色信息
func (m RoleApi) DropRoleByIds(ctx *gin.Context) {
	var ids []string
	err := ctx.ShouldBind(&ids)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	err = NewRoleApi().roleService.DropRole(ids)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, nil)
}

// 根据id获取角色信息
func (m RoleApi) GetRoleById(ctx *gin.Context) {
	id := ctx.Param("id")
	role, err := NewRoleApi().roleService.GetRoleById(id)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, role)
}
