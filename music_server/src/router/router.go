package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go_study/src/conf"
	"go_study/src/global"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

// IFnRegistRoute 定义注册路由的类型 rgPblicRoute公共路由,需要鉴权的路由
type IFnRegistRoute = func(rgPblicRoute *gin.RouterGroup, rgAuthRoute *gin.RouterGroup)

var (
	//gfnRouters 定义一个全局的路由数组
	GfnRouters []IFnRegistRoute
)

// InitRoute 初始化路由
func InitRoute() {
	//优雅退出web服务 context.Background()根上下文 syscall.SIGINT ctrl+c syscall.SIGTERM结束
	ctx, cancleCtx := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancleCtx()
	r := gin.Default()

	//跨域中间件
	r.Use(conf.Cors())

	//定义两个路由组
	rgPublic := r.Group("/api/v1/public")
	rgAuth := r.Group("/api/v1")

	//gin整合swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//路由鉴权
	//rgAuth.Use(middleware.JwtCheck())

	InitBasePlatformRoute()

	//将gfnRouters数组中的路由进行循环注册
	for _, fnRegistRoute := range GfnRouters {
		fnRegistRoute(rgPublic, rgAuth)
	}

	//读取配置文件中的端口号
	stPort := viper.Get("server.port")

	if stPort == "" {
		stPort = 8080
	}

	//创建httpserveer
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", stPort),
		Handler: r,
	}

	//启动web服务的监听
	go func() {
		global.Logger.Info(fmt.Sprintf("Start Server Listen Port :%s", stPort))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Error(fmt.Sprintf("Start Server Error:%s", err.Error()))
			return
		}
	}()
	<-ctx.Done()

	//创建一个带超时的上下文
	ctx, cancleShutdown := context.WithTimeout(context.Background(), time.Second*5)
	defer cancleShutdown()
	if err := server.Shutdown(ctx); err != nil {
		// TODO 记录日志
		global.Logger.Error(fmt.Sprintf("Stop Server Error : %s", err.Error()))
		return
	}
	global.Logger.Info("Stop Server Success")
}

// 初始化基础模块的路由
func InitBasePlatformRoute() {
	//初始化用户路由
	InitUserRoute()
	//初始化上传路由
	UploadFile()
	//初始歌手路由
	InitSingerRouter()
	//初始化菜单路由
	InitMenuRouter()
	//初始化字典类型路由
	InitDictTypeRouter()
	//初始化部门路由
	InitDepartmentRouter()
	//初始化角色路由
	InitRoleRouter()
	//初始化歌词路由
	InitLyricRouter()
	//初始化代码自动生成路由
	InitCodeGenRouter()

}

// RegistRoute 注册路由
func RegistRoute(fn IFnRegistRoute) {
	if fn == nil {
		return
	}
	GfnRouters = append(GfnRouters, fn)
}
