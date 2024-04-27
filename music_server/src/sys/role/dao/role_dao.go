package dao

import (
	"go_study/src/common"
	"go_study/src/sys/role/entity"
	"gorm.io/gorm"
)

type RoleDao struct {
	common.BaseDao
}

var roleDao *RoleDao

// NewRoleDao 做成单实列
func NewRoleDao() *RoleDao {
	if roleDao == nil {
		return &RoleDao{
			common.NewBaseDao(),
		}
	}
	return roleDao
}

// 获取角色列表
func (m RoleDao) GetRoleLists(role entity.Role, page common.Page) ([]entity.Role, int64, error) {
	var lists []entity.Role
	var count int64
	offset := (page.PageNum - 1) * page.PageSize
	tx := NewRoleDao().Orm.Where(role).Offset(offset).Limit(page.PageSize).Find(&lists).Offset(-1).Limit(-1).Count(&count)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	return lists, count, nil
}

// 添加角色信息
func (m RoleDao) AddRole(role entity.Role, tx *gorm.DB) error {
	err := tx.Create(&role)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

// 修改角色信息
func (m RoleDao) UpdateRole(role entity.Role, tx *gorm.DB) error {
	err := tx.Model(&entity.Role{}).Where("id = ?", role.ID).Updates(role)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

// 删除角色信息
func (m RoleDao) DropRole(id []string, tx *gorm.DB) error {
	var role entity.Role
	err := tx.Delete(&role, id)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

// 根据id获取角色信息
func (m RoleDao) GetRoleById(id string) (entity.Role, error) {
	var role entity.Role
	tx := NewRoleDao().Orm.Where("id = ?", id).First(&role)
	if tx.Error != nil {
		return role, tx.Error
	}
	return role, nil
}
