package router

import (
	"github.com/gin-gonic/gin"
	"go_study/src/sys/menu/api"
)

func InitMenuRouter() {
	//注册路由
	RegistRoute(func(rgPblicRoute *gin.RouterGroup, rgAuthRoute *gin.RouterGroup) {
		var menuApi = api.MenuApi{}
		group := rgAuthRoute.Group("menu")
		{
			//添加路由
			group.POST("/add", menuApi.AddMenu)
			//删除路由
			group.GET("/drop", menuApi.DropMenu)
			//修改路由
			group.POST("/update", menuApi.UpdateMenu)
			//查询路由
			group.POST("/list", menuApi.GetMenuLists)
			//添加路由时获取菜单树
			group.GET("/add/tree", menuApi.GetMenuAddTree)
			//根据id查询菜单
			group.GET("/byID", menuApi.GetMenuById)
		}
	})

}
