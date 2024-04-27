package dao

import (
	"errors"
	"go_study/src/common"
	"go_study/src/sys/menu/entity"
	"go_study/src/utils"
)

var menuDao *MenuDao

type MenuDao struct {
	common.BaseDao
}

// NewMenuDao 做成单实列
func NewMenuDao() *MenuDao {
	if menuDao == nil {
		return &MenuDao{
			common.NewBaseDao(),
		}
	}
	return menuDao

}

// 添加菜单
func (m MenuDao) AddMenu(menu entity.Menu) (int64, error) {
	menu.ID = utils.GenerateID()
	tx := NewMenuDao().Orm.Create(&menu)
	if tx.RowsAffected > 0 && tx.Error == nil {
		return tx.RowsAffected, nil
	} else {
		return 0, tx.Error
	}
}

// 删除菜单
func (m MenuDao) DropMenu(id string) error {
	var menu entity.Menu
	//删除菜单
	tx := NewMenuDao().Orm.Where("id = ?", id).Delete(&menu)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

// 修改菜单
func (m MenuDao) UpdateMenu(menu map[string]interface{}) error {
	id, ok := menu["id"]
	if !ok {
		return errors.New("参数中ID不存在")
	}
	tx := NewMenuDao().Orm.Table("tb_menu").Where("id = ?", id).Updates(menu)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

// 获取菜单列表
func (m MenuDao) GetMenuLists(queryMenu entity.Menu) ([]entity.Menu, error) {
	var menu []entity.Menu
	tx := NewMenuDao().Orm.Where(&queryMenu).Order("order_num").Find(&menu)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return menu, nil
}

// 根据id获取菜单
func (m MenuDao) GetMenuById(id string) (*entity.Menu, error) {
	var menu entity.Menu
	tx := NewMenuDao().Orm.Where("id = ?", id).First(&menu)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return &menu, nil
}
