package cmd

import (
	"fmt"
	"go_study/src/conf"
	"go_study/src/global"
	"go_study/src/router"
)

func Start() {
	//读取配置文件
	conf.InitConfig()
	//初始化日志文件
	global.Logger = conf.InitLogger()
	//初始化minio
	conf.InitMinio()
	//初始化数据库连接
	conf.InitDb()
	//初始化redis连接
	conf.InitRedis()
	//初始化路由 注：路由初始化必须放到最后
	router.InitRoute()
}

func Clear() {
	fmt.Println("退出程序")
}
