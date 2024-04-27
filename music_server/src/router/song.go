package router

import (
	"github.com/gin-gonic/gin"
	"go_study/src/core/song/api"
)

func UploadFile() {
	RegistRoute(func(rgPblicRoute *gin.RouterGroup, rgAuthRoute *gin.RouterGroup) {
		songApi := api.NewSongApi()
		rgAuth := rgAuthRoute.Group("song")
		{
			rgAuth.GET("/queryByName", songApi.QueryBySongName)
			rgAuth.POST("/addSong", songApi.AddSong)
			rgAuth.POST("/uploadSong", songApi.UploadSong)
			rgAuth.POST("/uploadSongCover", songApi.UploadSongCover)
			rgAuth.POST("/pageList", songApi.QuerySongListPage)
			rgAuth.GET("/dropUploadedFile", songApi.DropUploadedSong)
			rgAuth.GET("/dropSongById", songApi.DropSongById)
			rgAuth.POST("/update", songApi.UpdateSong)
		}
	})
}
