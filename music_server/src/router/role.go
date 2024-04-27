package router

import (
	"github.com/gin-gonic/gin"
	"go_study/src/sys/role/api"
)

func InitRoleRouter() {
	RegistRoute(func(rgPblicRoute *gin.RouterGroup, rgAuthRoute *gin.RouterGroup) {
		var roleApi = api.NewRoleApi()
		group := rgAuthRoute.Group("role")
		{
			//获取角色列表
			group.POST("/list", roleApi.GetRoleList)
			//添加角色信息
			group.POST("/add", roleApi.AddRole)
			//修改角色信息
			group.POST("/update", roleApi.UpdateRole)
			//删除角色信息
			group.POST("/drop", roleApi.DropRoleByIds)
			//根据id获取角色信息
			group.POST("/infoById", roleApi.GetRoleById)
		}
	})
}
