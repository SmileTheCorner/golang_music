package dao

import (
	"errors"
	"go_study/src/common"
	"go_study/src/core/song/entity"
	"go_study/src/utils"
	"gorm.io/gorm"
)

var songDao *SongDao

type SongDao struct {
	common.BaseDao
}

func NewSongDao() *SongDao {
	if songDao == nil {
		return &SongDao{
			common.NewBaseDao(),
		}
	}
	return songDao
}

// QueryBySongName 根据歌曲名查询歌曲
func (m SongDao) QueryBySongName(songName string) (int64, entity.Song, error) {
	var count int64
	var song entity.Song
	err := NewSongDao().Orm.Model(&song).Where("song_name = ?", songName).Count(&count).Error
	if err != nil {
		return -1, song, err
	}
	return count, song, err
}

// AddSong 添加歌曲
func (m SongDao) AddSong(song entity.Song) (error, int64) {
	tx := NewSongDao().Orm.Create(&song)
	return tx.Error, tx.RowsAffected
}

// DeleteSong 删除歌曲
func (m SongDao) DeleteSong(songId string) (error, int64) {
	var song entity.Song
	tx := NewSongDao().Orm.Where("id = ?", songId).Delete(&song)
	return tx.Error, tx.RowsAffected
}

// UpdateSong 修改歌曲
func (m SongDao) UpdateSong(songMap map[string]interface{}) (error, int64) {
	//获取Id
	id, ok := songMap["id"]
	if !ok {
		return errors.New("参数id有误"), 0
	}
	tx := common.NewBaseDao().Orm.Table("tb_song").Where("id", id).Updates(songMap)
	return tx.Error, tx.RowsAffected
}

// QuerySongById 根据ID查询歌曲
func (m SongDao) QuerySongById(songId uint) (entity.Song, error) {
	var song entity.Song
	tx := NewSongDao().Orm.Where("id = ?", songId).Take(&song)
	return song, tx.Error
}

// QuerySongListPage 分页查询歌曲列表
func (m SongDao) QuerySongListPage(page common.Page) ([]entity.Song, int64, error) {
	var song []entity.Song
	var count int64
	offset := (page.PageNum - 1) * page.PageSize
	if offset < 0 {
		offset = 0
	}
	var tx *gorm.DB
	word := page.KeyWord
	queryMap := make(map[string]interface{})
	if word != nil {
		m2 := word.(map[string]interface{})
		queryMap = utils.MapKeyConvUnderline(m2)
	}
	var nameLike string
	//循环map拼装查询参数
	for index, value := range queryMap {
		if index == "song_name" && value != "" {
			nameLike = "%" + value.(string) + "%"
			delete(queryMap, index)
		}
		switch value.(type) {
		case string:
			if value == "" {
				delete(queryMap, index)
			}
		case []interface{}:
			var array = value.([]interface{})
			for index, item := range array {
				if item == nil || item == "" {
					array = append(array[:index], array[index+1:]...)
				}
			}
			if len(array) == 0 {
				delete(queryMap, index)
			} else {
				queryMap[index] = array
			}
		}
	}
	if nameLike != "" {
		tx = NewSongDao().Orm.Where("song_name like ?", nameLike).Where(queryMap).Limit(page.PageSize).Offset(offset).Order("created_at desc").Find(&song).Offset(-1).Limit(-1).Count(&count)
	} else {
		tx = NewSongDao().Orm.Where(queryMap).Limit(page.PageSize).Offset(offset).Order("created_at desc").Find(&song).Offset(-1).Limit(-1).Count(&count)
	}
	return song, count, tx.Error
}
