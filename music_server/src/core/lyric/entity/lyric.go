package entity

import "go_study/src/common"

type Lyric struct {
	common.Model
	SongName string `gorm:"type:varchar(20);not null;comment:歌曲名称" json:"songName"`
	LyricUrl string `gorm:"type:varchar(255);not null;comment:歌词地址" json:"lyricUrl"`
}

func (Lyric) TableName() string {
	return "tb_lyric"
}
