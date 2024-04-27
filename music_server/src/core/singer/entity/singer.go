package entity

import "go_study/src/common"

type Singer struct {
	common.Model
	SingerName        string `gorm:"type:varchar(20);not null;comment:歌手名字" json:"singerName"`
	SingerEHName      string `gorm:"type:varchar(20);comment:歌手英文名字" json:"singerEHName"`
	SingerAvatar      string `gorm:"type:varchar(100);comment:歌手头像Url地址" json:"singerAvatar"`
	SingerDescription string `gorm:"type:varchar(255);comment:歌手描述" json:"singerDescription"`
	SingerClassify    int8   `gorm:"type:tinyint(1);comment:歌手分类" json:"singerClassify"`
}
