package entity

import "go_study/src/common"

type DictType struct {
	common.Model
	ParentId string     `gorm:"type:varchar(200);comment:父id" json:"parentId"`
	DictName string     `gorm:"type:varchar(60);not null;comment:字典名称" json:"dictName"`
	DictType string     `gorm:"type:varchar(60);not null;comment:字典类型" json:"dictType"`
	Status   int8       `gorm:"type:tinyint(1);default:0;comment:字典状态（0正常 1停用）" json:"status"`
	Remark   string     `gorm:"type:varchar(255);comment:备注" json:"remark"`
	Children []DictType `gorm:"-" json:"children"`
}

func (DictType) TableName() string {
	return "tb_dict_type"
}
