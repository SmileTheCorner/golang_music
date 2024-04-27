package router

import (
	"github.com/gin-gonic/gin"
	"go_study/src/core/singer/api"
)

func InitSingerRouter() {
	RegistRoute(func(rgPblicRoute *gin.RouterGroup, rgAuthRoute *gin.RouterGroup) {
		singerApi := api.SingerApi{}
		rgSingerAuth := rgAuthRoute.Group("singer")
		{
			//新增歌手信息
			rgSingerAuth.POST("/add", singerApi.AddSinger)
			//查询歌手列表
			rgSingerAuth.POST("/queryList", singerApi.QuerySingerList)
			//根据歌手名查询歌手
			rgSingerAuth.GET("/querySingerByName/:name", singerApi.QuerySingerByName)
			//上传歌手头像
			rgSingerAuth.POST("/uploadSingerAvatar", singerApi.UploadSingerAvatar)
			//编辑歌手信息
			rgSingerAuth.POST("/edit", singerApi.EditSinger)
			//根据id切片删除歌手信息
			rgSingerAuth.DELETE("/dropByIds", singerApi.DeleteSingerByIds)
		}
	})
}
