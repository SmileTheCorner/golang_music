package entity

import "go_study/src/common"

type DictData struct {
	common.Model
	DictSort  int    `gorm:"default:0;comment:字典排序" json:"sort"`
	DictLabel string `gorm:"type:varchar(60);comment:字典标签" json:"label"`
	DictValue string `gorm:"type:varchar(60);comment:字典值" json:"value"`
	DictType  string `gorm:"type:varchar(60);comment:字典类型" json:"dictType"`
	Status    string `gorm:"type:char(1);default:0;comment:状态（0正常 1停用）" json:"status"`
	Remark    string `gorm:"type:varchar(255);comment:备注" json:"remark"`
}

func (DictData) TableName() string {
	return "tb_dict_data"
}
