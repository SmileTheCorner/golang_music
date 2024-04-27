package entity

import (
	"go_study/src/common"
)

// Menu 菜单表
type Menu struct {
	common.Model
	Name      string `gorm:"type:varchar(20);not null;comment:菜单名称" json:"name"`
	Title     string `gorm:"type:varchar(20);not null;comment:标题" json:"title"`
	ParentId  string `gorm:"type:varchar(200);comment:父ID" json:"parentId"`
	OrderNum  int8   `gorm:"type:tinyint(2);comment:显示顺序" json:"orderNum"`
	Path      string `gorm:"type:varchar(200);comment:路由地址" json:"path"`
	Redirect  string `gorm:"type:varchar(200);comment:重定向地址" json:"redirect"`
	Component string `gorm:"type:varchar(200);comment:组件路径" json:"component"`
	Query     string `gorm:"type:varchar(60);comment:路由参数" json:"query"`
	IsFrame   int8   `gorm:"type:tinyint(1);default:1;comment:是否为外链（0是 1否）" json:"isFrame"`
	IsCache   int8   `gorm:"type:tinyint(1);default:0;comment:是否缓存（0缓存 1不缓存）" json:"isCache"`
	IsFull    int8   `gorm:"type:tinyint(1);default:1;comment:是否全屏（0不全屏 1全屏）" json:"isFull"`
	IsLink    string `gorm:"type:varchar(200);comment:外链地址" json:"isLink"`
	MenuType  string `gorm:"type:varchar(1);comment:M目录 C菜单 F按钮）" json:"menuType"`
	Visible   int8   `gorm:"type:tinyint(1);default:0;comment:显示状态（0显示 1隐藏）" json:"visible"`
	Status    int8   `gorm:"type:tinyint(1);default:0;comment:菜单状态（0正常 1停用）" json:"status"`
	Perms     string `gorm:"type:varchar(200);comment:权限字符串" json:"perms"`
	Icon      string `gorm:"type:varchar(20);comment:菜单图标" json:"icon"`
	Remark    string `gorm:"type:varchar(200);comment:备注信息" json:"remark"`
	Children  []Menu `gorm:"-" json:"children"`
}

func (Menu) TableName() string {
	return "tb_menu"
}
