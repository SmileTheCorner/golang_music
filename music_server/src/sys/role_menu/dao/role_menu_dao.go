package dao

import (
	"go_study/src/common"
	"go_study/src/sys/role_menu/entity"
	roleMenuEntity "go_study/src/sys/role_menu/entity"
	"gorm.io/gorm"
)

type RoleMenuDao struct {
	common.BaseDao
}

var roleMenuDao *RoleMenuDao

// NewRoleMenuDao 做成单实列
func NewRoleMenuDao() *RoleMenuDao {
	if roleMenuDao == nil {
		return &RoleMenuDao{
			common.NewBaseDao(),
		}
	}
	return roleMenuDao
}

// 删除角色菜单信息
func (m RoleMenuDao) DropRoleMenu(roleId []string, tx *gorm.DB) error {
	var roleMenu entity.RoleMenu
	err := tx.Where("role_id = ?", roleId).Delete(&roleMenu)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

// 添加角色菜单信息
func (m RoleMenuDao) AddRoleMenu(roleMenu []entity.RoleMenu, tx *gorm.DB) error {
	err := tx.Create(&roleMenu)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

// 获取角色菜单信息
func (m RoleMenuDao) GetRoleMenu(roleId string) ([]entity.RoleMenu, error) {
	var roleMenu []entity.RoleMenu
	tx := NewRoleMenuDao().Orm.Where("role_id = ?", roleId).Find(&roleMenu)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return roleMenu, nil
}

// 修改角色菜单信息
func (m RoleMenuDao) UpdateRoleMenu(roleMenuList []roleMenuEntity.RoleMenu, tx *gorm.DB) error {
	//先根据roleId删除角色菜单信息
	err := tx.Where("role_id = ?", roleMenuList[0].RoleId).Delete(&roleMenuEntity.RoleMenu{})
	if err.Error != nil {
		return err.Error
	}
	createErr := tx.Create(&roleMenuList)
	if createErr.Error != nil {
		return createErr.Error
	}
	return nil
}
