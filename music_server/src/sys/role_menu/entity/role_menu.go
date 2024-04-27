package entity

type RoleMenu struct {
	RoleId string `gorm:"type:varchar(200);not null;comment:角色ID" json:"roleId"`
	MenuId string `gorm:"type:varchar(200);not null;comment:菜单ID" json:"menuId"`
}
