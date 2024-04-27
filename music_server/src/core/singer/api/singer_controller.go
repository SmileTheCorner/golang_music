package api

import (
	"github.com/gin-gonic/gin"
	"go_study/src/common"
	"go_study/src/conf"
	"go_study/src/core/singer/entity"
	"go_study/src/core/singer/service"
	"go_study/src/enum"
	"go_study/src/utils"
)

type SingerApi struct {
	singerService *service.SingerService
}

func NewSingerApi() SingerApi {
	return SingerApi{
		singerService: service.NewSingerService(),
	}
}

// AddSinger 新增歌手信息
func (m SingerApi) AddSinger(ctx *gin.Context) {
	var singer entity.Singer
	err := ctx.ShouldBind(&singer)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	id, adderr := NewSingerApi().singerService.AddSinger(singer)
	if adderr != nil {
		utils.Fail(ctx, adderr)
		return
	}
	utils.Ok(ctx, id)
}

// QuerySingerList 查询歌手列表
func (m SingerApi) QuerySingerList(ctx *gin.Context) {
	//获取请求体中的参数
	var page common.Page
	err := ctx.ShouldBind(&page)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	singerList, err, count := NewSingerApi().singerService.QuerySingerList(page)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	//进行数据的封装
	mapData := utils.PageData(singerList, int64(page.PageNum), int64(page.PageSize), count)
	utils.Ok(ctx, mapData)
}

// QuerySingerByName 根据歌手名查询用户
func (m SingerApi) QuerySingerByName(ctx *gin.Context) {
	singerName := ctx.Param("name")
	count, _ := NewSingerApi().singerService.QuerySingerByName(singerName)
	utils.Ok(ctx, count)
}

// UploadSingerAvatar 上传歌手头像
func (m SingerApi) UploadSingerAvatar(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	str, err := conf.UploadFile(file, enum.BucketName, enum.SingerAvatar)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, str)
}

// 编辑歌手信息
func (m SingerApi) EditSinger(ctx *gin.Context) {
	var singer entity.Singer
	err := ctx.ShouldBind(&singer)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	err = NewSingerApi().singerService.EditSinger(singer)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, "")
}

// 根据Id切片删除歌手信息
func (m SingerApi) DeleteSingerByIds(ctx *gin.Context) {
	var ids []string
	err := ctx.ShouldBind(&ids)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	err = NewSingerApi().singerService.DeleteSingerByIds(ids)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, "")
}
