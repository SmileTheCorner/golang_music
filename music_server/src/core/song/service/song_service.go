package service

import (
	"go_study/src/common"
	"go_study/src/core/song/dao"
	"go_study/src/core/song/entity"
	"go_study/src/utils"
	"time"
)

var songService *SongService

type SongService struct {
	common.BaseService
	songDao *dao.SongDao
}

func NewSongService() *SongService {
	if songService == nil {
		return &SongService{
			songDao: dao.NewSongDao(),
		}
	}
	return songService
}

// QueryBySongName 根据歌曲名查询歌曲
func (m SongService) QueryBySongName(songName string) (int64, entity.Song, error) {
	count, song, err := NewSongService().songDao.QueryBySongName(songName)
	return count, song, err
}

// DeleteSong 删除歌曲
func (m SongService) DeleteSong(songId string) (error, int64) {
	return NewSongService().songDao.DeleteSong(songId)
}

// UpdateSong 修改歌曲
func (m SongService) UpdateSong(songMap map[string]interface{}) (error, int64) {
	now := time.Now()
	format := now.Format(time.DateTime)
	songMap["updatedAt"] = format
	songMap = utils.MapKeyConvUnderline(songMap)
	err, i := NewSongService().songDao.UpdateSong(songMap)
	return err, i
}

// QuerySongListPage 分页查询歌曲列表
func (m SongService) QuerySongListPage(page common.Page) ([]entity.Song, int64, error) {
	listPage, count, err := NewSongService().songDao.QuerySongListPage(page)
	if err != nil {
		return nil, 0, err
	}
	return listPage, count, nil
}

// DropSongById 根据ID删除歌曲信息
func (m SongService) DropSongById(id string) (error, int64) {
	err, row := NewSongService().songDao.DeleteSong(id)
	return err, row
}
