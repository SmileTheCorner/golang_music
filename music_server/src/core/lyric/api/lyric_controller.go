package api

import (
	"github.com/gin-gonic/gin"
	"go_study/src/common"
	"go_study/src/core/lyric/entity"
	"go_study/src/core/lyric/service"
	"go_study/src/utils"
	"strconv"
)

type LyricApi struct {
	lyricService *service.LyricService
}

func NewLyricApi() LyricApi {
	return LyricApi{
		lyricService: service.NewLyricService(),
	}
}

// UploadLyric 上传歌词
func (m LyricApi) UploadLyric(ctx *gin.Context) {
	//从请求体中获取到要上传的文件的二进制流
	file, err := ctx.FormFile("file")
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	lyric, err := NewLyricApi().lyricService.UploadLyric(file)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, lyric)
}

// AddLyric 添加歌词
func (m LyricApi) AddLyric(ctx *gin.Context) {
	var lyric entity.Lyric
	err := ctx.ShouldBind(&lyric)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	addLyric, err := NewLyricApi().lyricService.AddLyric(lyric)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, addLyric)
}

// DeleteLyric 删除歌词
func (m LyricApi) DeleteLyric(ctx *gin.Context) {
	var ids []string
	paramsErr := ctx.ShouldBind(&ids)
	if paramsErr != nil {
		utils.Fail(ctx, paramsErr)
		return
	}
	lyric, err := NewLyricApi().lyricService.DeleteLyric(ids)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, lyric)
}

// FindLyricByID 根据ID查询歌词信息
func (m LyricApi) FindLyricByID(ctx *gin.Context) {
	value := ctx.Query("lyricId")
	lyricId, _ := strconv.Atoi(value)
	lyric, err := NewLyricApi().lyricService.FindLyricByID(uint(lyricId))
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, lyric)
}

// PageFindLyricList 分页查询歌词列表
func (m LyricApi) PageFindLyricList(ctx *gin.Context) {
	//获取请求体中的参数
	var page common.Page
	err := ctx.ShouldBind(&page)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	list, err, count := NewLyricApi().lyricService.PageFindLyricList(page)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	//进行数据的封装
	mapData := utils.PageData(list, int64(page.PageNum), int64(page.PageSize), count)
	utils.Ok(ctx, mapData)
}

// 修改歌词
func (m LyricApi) UpdateLyric(ctx *gin.Context) {
	var lyric entity.Lyric
	err := ctx.ShouldBind(&lyric)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	err = NewLyricApi().lyricService.UpdateLyric(lyric)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, err)
}
