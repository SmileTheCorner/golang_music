package dao

import (
	"go_study/src/common"
	"go_study/src/core/singer/entity"
	"go_study/src/utils"
)

var singerDao *SingerDao

type SingerDao struct {
	common.BaseDao
}

func NewSingerDao() *SingerDao {
	if singerDao == nil {
		return &SingerDao{
			common.NewBaseDao(),
		}
	}
	return singerDao
}

// AddSinger 新增歌手信息
func (m SingerDao) AddSinger(singer entity.Singer) (string, error) {
	id := utils.GenerateID()
	singer.ID = id
	result := m.Orm.Create(&singer)
	if err := result.Error; err != nil {
		return "", err
	}
	return singer.ID, nil
}

// QuerySingerList 查询歌手列表
func (m SingerDao) QuerySingerList(page common.Page) ([]entity.Singer, error, int64) {
	var singerList []entity.Singer
	var count int64

	queryMap := utils.PageToObject(&page)
	tx := m.Orm.Where(queryMap).Limit(page.PageSize).Offset(page.Offset).Find(&singerList).Offset(-1).Limit(-1).Count(&count)

	return singerList, tx.Error, count
}

// QuerySingerByName 根据歌手名查询歌手
func (m SingerDao) QuerySingerByName(singerName string) (int64, entity.Singer) {
	var singer entity.Singer
	var count int64
	m.Orm.Model(&singer).Where("singer_name = ?", singerName).Count(&count)
	return count, singer
}

// 编辑歌手信息
func (m SingerDao) EditSinger(singer entity.Singer) error {
	result := m.Orm.Save(&singer)
	return result.Error
}

// 根据id切片删除歌手信息
func (m SingerDao) DeleteSingerByIds(ids []string) error {
	result := m.Orm.Delete(&entity.Singer{}, ids)
	return result.Error
}
