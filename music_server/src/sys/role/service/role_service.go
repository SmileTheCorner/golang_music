package service

import (
	"encoding/json"
	"go_study/src/common"
	"go_study/src/sys/role/dao"
	"go_study/src/sys/role/entity"
	roleMenuDao "go_study/src/sys/role_menu/dao"
	roleMenuEntity "go_study/src/sys/role_menu/entity"
	"go_study/src/utils"
)

type RoleService struct {
	roleDao     *dao.RoleDao
	roleMenuDao *roleMenuDao.RoleMenuDao
	common.BaseService
}

var roleService *RoleService

// NewRoleService 做成单实列
func NewRoleService() *RoleService {
	if roleService == nil {
		return &RoleService{
			roleDao:     dao.NewRoleDao(),
			roleMenuDao: roleMenuDao.NewRoleMenuDao(),
		}
	}
	return roleService
}

// 获取角色列表
func (m RoleService) GetRoleLists(role entity.Role, page common.Page) ([]entity.Role, int64, error) {
	roleLists, count, err := NewRoleService().roleDao.GetRoleLists(role, page)
	if err != nil {
		return nil, 0, err
	}
	return roleLists, count, nil
}

// 添加角色信息
func (m RoleService) AddRole(roleReq entity.RoleReq) error {
	role := entity.Role{}
	data, jsonErr := json.Marshal(roleReq.DataScopeDept)
	if jsonErr != nil {
		return jsonErr
	}
	id := utils.GenerateID()
	role.ID = id
	role.Name = roleReq.Name
	role.Code = roleReq.Code
	role.Sort = roleReq.Sort
	role.DataScope = roleReq.DataScope
	role.DataScopeDept = string(data)
	role.Status = roleReq.Status
	role.Remark = roleReq.Remark

	//获取到菜单切片
	var menuList = roleReq.MenuList
	var roleMenuList []roleMenuEntity.RoleMenu
	for _, item := range menuList {
		var roleMenu roleMenuEntity.RoleMenu
		roleMenu.RoleId = id
		roleMenu.MenuId = item
		roleMenuList = append(roleMenuList, roleMenu)
	}

	//开启事物
	tx := common.NewBaseDao().Orm.Begin()
	//抛出异常进行事物回滚
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			return
		}
	}()
	if tx.Error != nil {
		return tx.Error
	}

	err := NewRoleService().roleDao.AddRole(role, tx)
	if err != nil {
		//事物回滚
		tx.Rollback()
		return err
	}

	roleMenuErr := NewRoleService().roleMenuDao.AddRoleMenu(roleMenuList, tx)
	if roleMenuErr != nil {
		//事物回滚
		tx.Rollback()
		return roleMenuErr
	}
	//提交事物
	tx.Commit()
	return nil
}

// 修改角色信息
func (m RoleService) UpdateRole(roleReq entity.RoleReq) error {
	role := entity.Role{}
	if len(roleReq.DataScopeDept) > 0 {
		data, jsonErr := json.Marshal(roleReq.DataScopeDept)
		if jsonErr != nil {
			return jsonErr
		}
		role.DataScopeDept = string(data)
	}

	role.ID = roleReq.ID
	role.Name = roleReq.Name
	role.Code = roleReq.Code
	role.Sort = roleReq.Sort
	role.DataScope = roleReq.DataScope
	role.Status = roleReq.Status
	role.Remark = roleReq.Remark

	//获取到菜单切片
	var menuList = roleReq.MenuList
	var roleMenuList []roleMenuEntity.RoleMenu
	for _, item := range menuList {
		var roleMenu roleMenuEntity.RoleMenu
		roleMenu.RoleId = roleReq.ID
		roleMenu.MenuId = item
		roleMenuList = append(roleMenuList, roleMenu)
	}

	//开启事物
	tx := common.NewBaseDao().Orm.Begin()
	//抛出异常进行事物回滚
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			return
		}
	}()
	if tx.Error != nil {
		return tx.Error
	}

	err := NewRoleService().roleDao.UpdateRole(role, tx)
	if err != nil {
		//事物回滚
		tx.Rollback()
		return err
	}

	roleMenuErr := NewRoleService().roleMenuDao.UpdateRoleMenu(roleMenuList, tx)
	if roleMenuErr != nil {
		//事物回滚
		tx.Rollback()
		return roleMenuErr
	}

	//提交事务
	tx.Commit()

	return nil
}

// 删除角色信息
func (m RoleService) DropRole(id []string) error {
	//开启事务
	tx := common.NewBaseDao().Orm.Begin()
	//抛出异常进行事物回滚
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			return
		}
	}()
	if tx.Error != nil {
		return tx.Error
	}
	err := NewRoleService().roleDao.DropRole(id, tx)
	if err != nil {
		//事物回滚
		tx.Rollback()
		return err
	}
	NewRoleService().roleMenuDao.DropRoleMenu(id, tx)
	//提交事务
	tx.Commit()
	return nil
}

// 根据id获取角色信息
func (m RoleService) GetRoleById(id string) (entity.Role, error) {
	role, err := NewRoleService().roleDao.GetRoleById(id)
	if err != nil {
		return role, err
	}
	return role, nil
}
