package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_study/src/common"
	"go_study/src/conf"
	"go_study/src/core/song/entity"
	"go_study/src/core/song/service"
	"go_study/src/enum"
	"go_study/src/global"
	"go_study/src/utils"
)

var songApiInstance *songApi

type songApi struct {
	songService *service.SongService
}

func NewSongApi() *songApi {
	if songApiInstance == nil {
		return &songApi{
			songService: service.NewSongService(),
		}
	}
	return songApiInstance
}

// QueryBySongName 根据歌曲名查询歌曲
func (m songApi) QueryBySongName(ctx *gin.Context) {
	//获取歌曲传过来的参数
	songName := ctx.Query("name")
	count, _, err := NewSongApi().songService.QueryBySongName(songName)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, count)
}

// AddSong 添加歌曲信息
func (m songApi) AddSong(ctx *gin.Context) {
	//获取传过来的参数，使用结构体接收
	var song entity.Song
	ctx.ShouldBind(&song)
	//将数据添加到数据库中
	song.ID = utils.GenerateID()
	db := global.DB.Create(&song)
	rowAffect := db.RowsAffected
	err := db.Error
	if rowAffect > 0 && err == nil {
		utils.Ok(ctx, rowAffect)
	} else {
		utils.Fail(ctx, err)
	}
}

// UpdateSong 更新歌曲信息
func (m songApi) UpdateSong(ctx *gin.Context) {
	//获取需要跟新歌曲的信息
	songMap := make(map[string]interface{})
	if err := ctx.BindJSON(&songMap); err != nil {
		utils.Fail(ctx, err)
		return
	}
	err, row := NewSongApi().songService.UpdateSong(songMap)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, row)
}

// UploadSong 上传歌曲
func (m songApi) UploadSong(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	str, err := conf.UploadFile(file, enum.BucketName, enum.Song)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, str)
}

// UploadSongCover 上传歌曲封面
func (m songApi) UploadSongCover(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	str, err := conf.UploadFile(file, enum.BucketName, enum.SongCover)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, str)
}

// QuerySongListPage 分页查询歌曲列表
func (m songApi) QuerySongListPage(ctx *gin.Context) {
	var page common.Page
	ctx.ShouldBind(&page)
	listPage, count, err := NewSongApi().songService.QuerySongListPage(page)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	//进行数据的封装
	mapData := utils.PageData(listPage, int64(page.PageNum), int64(page.PageSize), count)
	utils.Ok(ctx, mapData)
}

// DropUploadedSong 删除上传成功的歌曲
func (m songApi) DropUploadedSong(ctx *gin.Context) {
	fileName, b := ctx.GetQuery("fileName")
	if !b {
		fmt.Println("参数传递错误")
		return
	}
	err := conf.RemoveObject(enum.BucketName, fileName)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, "删除成功")
}

// DropSongById 根据ID删除歌曲信息
func (m songApi) DropSongById(ctx *gin.Context) {
	//获取参数
	songID, b := ctx.GetQuery("songId")
	if !b {
		utils.Fail(ctx, errors.New("参数传递错误"))
		return
	}
	err, row := NewSongApi().songService.DropSongById(songID)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, row)
}
