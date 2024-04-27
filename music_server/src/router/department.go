package router

import (
	"github.com/gin-gonic/gin"
	"go_study/src/sys/department/api"
)

func InitDepartmentRouter() {
	RegistRoute(func(rgPblicRoute *gin.RouterGroup, rgAuthRoute *gin.RouterGroup) {
		var departmentApi = api.NewDepartmentApi()
		group := rgAuthRoute.Group("department")
		{
			//获取部门列表
			group.POST("/list", departmentApi.GetDepartmentList)
			//添加部门信息
			group.POST("/add", departmentApi.AddDepartment)
			//修改部门信息
			group.POST("/update", departmentApi.UpdateDepartment)
			//删除部门信息
			group.POST("/drop", departmentApi.DropDepartment)
			//获取部门树形结构
			group.POST("/tree", departmentApi.GetDepartmentTree)
		}
	})
}
