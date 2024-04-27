package router

import (
	"github.com/gin-gonic/gin"
	"go_study/src/sys/user/api"
)

func InitUserRoute() {
	RegistRoute(func(rgPblicRoute *gin.RouterGroup, rgAuthRoute *gin.RouterGroup) {
		userApi := api.NewUserApi()
		//公共登录
		rgPublicUser := rgPblicRoute.Group("user")
		{
			rgPublicUser.POST("/login", userApi.Login)
		}

		//需要鉴权的用户组
		rgUserAuth := rgAuthRoute.Group("user")
		{
			//根据用户名查找用户
			rgUserAuth.GET("/queryUserByName", userApi.QueryUserByUserName)
			//添加新用户
			rgUserAuth.POST("/addUser", userApi.AddUser)
			//获取用户列表
			rgUserAuth.POST("/list", userApi.QueryUserList)
		}

	})
}
