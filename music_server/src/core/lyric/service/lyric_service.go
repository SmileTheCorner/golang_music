package service

import (
	"go_study/src/common"
	"go_study/src/core/lyric/dao"
	"go_study/src/core/lyric/entity"
	"go_study/src/utils"
	"mime/multipart"
)

var lyricService *LyricService

type LyricService struct {
	common.BaseService
	lyricDao *dao.LyricDao
}

func NewLyricService() *LyricService {
	if lyricService == nil {
		return &LyricService{
			lyricDao: dao.NewLyricDao(),
		}
	} else {
		return lyricService
	}
}

// UploadLyric 上传歌词
func (m LyricService) UploadLyric(lyric *multipart.FileHeader) (string, error) {
	uploadLyric, err := NewLyricService().lyricDao.UploadLyric(lyric)
	if err != nil {
		return "", err
	}
	return uploadLyric, nil
}

// AddLyric 添加歌词
func (m LyricService) AddLyric(lyric entity.Lyric) (int64, error) {
	lyric.ID = utils.GenerateID()
	addLyric, err := NewLyricService().lyricDao.AddLyric(lyric)
	if err != nil {
		return 0, err
	}
	return addLyric, nil
}

// DeleteLyric 删除歌词
func (m LyricService) DeleteLyric(ids []string) (int64, error) {
	id, err := NewLyricService().lyricDao.DeleteLyricById(ids)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// FindLyricByID 根据ID查询歌词信息
func (m LyricService) FindLyricByID(lyricId uint) (entity.Lyric, error) {
	lyric, err := NewLyricService().lyricDao.FindLyricByID(lyricId)
	if err != nil {
		return lyric, err
	}
	return lyric, nil
}

// PageFindLyricList 分页查询歌词列表
func (m LyricService) PageFindLyricList(page common.Page) ([]entity.Lyric, error, int64) {
	list, err, count := NewLyricService().lyricDao.PageFindLyricList(page)
	return list, err, count
}

// 修改歌词
func (m LyricService) UpdateLyric(lyric entity.Lyric) error {
	err := NewLyricService().lyricDao.UpdateLyric(lyric)
	return err
}
