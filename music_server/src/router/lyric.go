package router

import (
	"github.com/gin-gonic/gin"
	"go_study/src/core/lyric/api"
)

func InitLyricRouter() {
	RegistRoute(func(rgPblicRoute *gin.RouterGroup, rgAuthRoute *gin.RouterGroup) {
		lyricApi := api.LyricApi{}
		group := rgAuthRoute.Group("/lyric")
		{
			//获取歌词列表
			group.POST("/list", lyricApi.PageFindLyricList)
			//添加歌词
			group.POST("/add", lyricApi.AddLyric)
			//上传歌词文件
			group.POST("/upload", lyricApi.UploadLyric)
			//删除歌词
			group.DELETE("/drop", lyricApi.DeleteLyric)
			//根据id查询歌词信息
			group.GET("/queryById", lyricApi.FindLyricByID)
			//修改歌词
			group.PUT("/edit", lyricApi.UpdateLyric)
		}
	})
}
