package dao

import (
	"go_study/src/common"
	"go_study/src/conf"
	"go_study/src/core/lyric/entity"
	"go_study/src/core/song/dao"
	"go_study/src/enum"
	"go_study/src/utils"
	"mime/multipart"
)

var lyricDao *LyricDao

type LyricDao struct {
	common.BaseDao
}

func NewLyricDao() *LyricDao {
	if lyricDao == nil {
		return &LyricDao{
			common.NewBaseDao(),
		}
	}
	return lyricDao
}

// UploadLyric 上传歌词
func (m LyricDao) UploadLyric(lyric *multipart.FileHeader) (string, error) {
	fileUrl, err := conf.UploadFile(lyric, enum.BucketName, enum.Lyric)
	return fileUrl, err
}

// AddLyric 添加歌词
func (m LyricDao) AddLyric(lyric entity.Lyric) (int64, error) {
	tx := NewLyricDao().Orm.Create(&lyric)
	return tx.RowsAffected, tx.Error
}

// DeleteLyricById 根据ID删除歌词
func (m LyricDao) DeleteLyricById(ids []string) (int64, error) {
	var lyric entity.Lyric
	tx := NewLyricDao().Orm.Delete(&lyric, ids)
	return tx.RowsAffected, tx.Error
}

// FindLyricByID 根据ID查询歌词信息
func (m LyricDao) FindLyricByID(lyricId uint) (entity.Lyric, error) {
	var lyric entity.Lyric
	tx := NewLyricDao().Orm.Where("id=?", lyricId).First(&lyric)
	if tx.Error != nil {
		return lyric, tx.Error
	}
	return lyric, nil
}

// PageFindLyricList 分页查询歌词列表
func (m LyricDao) PageFindLyricList(page common.Page) ([]entity.Lyric, error, int64) {
	var lyric []entity.Lyric
	var count int64
	queryMap := utils.PageToObject(&page)
	tx := dao.NewSongDao().Orm.Where(queryMap).Limit(page.PageSize).Offset(page.Offset).Find(&lyric).Offset(-1).Limit(-1).Count(&count)
	return lyric, tx.Error, count
}

// 修改歌词
func (m LyricDao) UpdateLyric(lyric entity.Lyric) error {
	tx := NewLyricDao().Orm.Model(&lyric).Updates(lyric)
	return tx.Error
}
