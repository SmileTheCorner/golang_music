package entity

import "go_study/src/common"

type Department struct {
	common.Model
	ParentId string       `gorm:"type:varchar(200);comment:父ID" json:"parentId"`
	Name     string       `gorm:"type:varchar(20);not null;comment:部门名称" json:"name"`
	Contacts string       `gorm:"type:varchar(20);comment:部门联系人" json:"contacts"`
	Phone    string       `gorm:"type:varchar(11);comment:联系人电话" json:"phone"`
	Email    string       `gorm:"type:varchar(20);comment:联系人邮箱" json:"email"`
	OrderNum int8         `gorm:"type:tinyint(2);comment:显示顺序" json:"orderNum"`
	Status   int8         `gorm:"type:tinyint(1);default:0;comment:部门状态（0正常 1停用）" json:"status"`
	Children []Department `gorm:"-" json:"children"`
}
