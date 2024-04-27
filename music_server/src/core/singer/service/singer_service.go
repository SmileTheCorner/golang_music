package service

import (
	"go_study/src/common"
	"go_study/src/core/singer/dao"
	"go_study/src/core/singer/entity"
)

var singerService *SingerService

type SingerService struct {
	common.BaseService
	SingerDao *dao.SingerDao
}

func NewSingerService() *SingerService {
	if singerService == nil {
		return &SingerService{
			SingerDao: dao.NewSingerDao(),
		}
	}
	return singerService
}

// AddSinger 新增歌手信息
func (m SingerService) AddSinger(singer entity.Singer) (string, error) {
	id, err := NewSingerService().SingerDao.AddSinger(singer)
	return id, err
}

// QuerySingerList 查询歌手列表
func (m SingerService) QuerySingerList(page common.Page) ([]entity.Singer, error, int64) {
	singerList, err, count := NewSingerService().SingerDao.QuerySingerList(page)
	return singerList, err, count
}

// QuerySingerByName 根据歌手名查询用户
func (m SingerService) QuerySingerByName(singerName string) (int64, entity.Singer) {
	count, singer := NewSingerService().SingerDao.QuerySingerByName(singerName)
	return count, singer
}

// 编辑歌手信息
func (m SingerService) EditSinger(singer entity.Singer) error {
	err := NewSingerService().SingerDao.EditSinger(singer)
	return err
}

// 根据id切片删除歌手信息
func (m SingerService) DeleteSingerByIds(ids []string) error {
	err := NewSingerService().SingerDao.DeleteSingerByIds(ids)
	return err
}
