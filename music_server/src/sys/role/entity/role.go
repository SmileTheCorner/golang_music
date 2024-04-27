package entity

import "go_study/src/common"

type Role struct {
	common.Model
	Name          string `gorm:"type:varchar(20);comment:角色名称;not null" json:"name"`
	Code          string `gorm:"type:varchar(20);comment:角色权限字符串;not null" json:"code"`
	Sort          int8   `gorm:"type:tinyint;comment:显示顺序" json:"sort"`
	DataScope     int8   `gorm:"type:tinyint;comment:数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）" json:"dataScope"`
	DataScopeDept string `gorm:"type:varchar(225);comment:数据范围(指定部门数组)" json:"dataScopeDept"`
	Status        int8   `gorm:"type:tinyint;comment:角色状态（0正常 1停用）" json:"status"`
	Remark        string `gorm:"type:varchar(500);comment:备注" json:"remark"`
}

type RoleReq struct {
	common.Model
	Name          string   `json:"name" binding:"required"`
	Code          string   `json:"code" binding:"required"`
	Sort          int8     `json:"sort"`
	MenuList      []string `json:"menuList"`
	DataScope     int8     `json:"dataScope"`
	DataScopeDept []string `json:"dataScopeDept"`
	Status        int8     `json:"status"`
	Remark        string   `json:"remark"`
}
