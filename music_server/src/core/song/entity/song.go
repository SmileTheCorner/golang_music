package entity

import "go_study/src/common"

type Song struct {
	common.Model
	SongName        string `gorm:"size:255" json:"songName"`                  //歌曲名称
	SongUrl         string `gorm:"size:255" json:"songUrl"`                   //歌曲地址
	SongCover       string `gorm:"size:255" json:"songCover"`                 //歌曲封面
	SongSingerId    string `gorm:"column:song_singer_id" json:"songSingerId"` //歌手id
	LyricsId        string `gorm:"column:lyrics_id" json:"lyricsId"`          //歌词id
	SongDescription string `gorm:"size:255" json:"songDescription"`           //歌曲描述
	SongType        string `json:"songType"`                                  //歌曲类型
	Recommend       bool   `json:"recommend"`                                 //是否推荐
	IsNew           bool   `json:"isNew"`                                     //是否是新音乐
	RankingListId   string `json:"rankingListId"`                             //排行榜id
	PlayCount       uint64 `json:"playCount"`                                 //播放次数
}
