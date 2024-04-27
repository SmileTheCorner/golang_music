package service

import (
	"go_study/src/common"
	"go_study/src/sys/menu/dao"
	"go_study/src/sys/menu/entity"
	"go_study/src/utils"
)

var menuService *MenuService

type MenuService struct {
	common.BaseService
	menuDao *dao.MenuDao
}

func NewMenuService() *MenuService {
	if menuService == nil {
		return &MenuService{
			menuDao: dao.NewMenuDao(),
		}
	}
	return menuService
}

// 添加菜单
func (m MenuService) AddMenu(menu entity.Menu) (int64, error) {
	//添加menu
	rows, err := NewMenuService().menuDao.AddMenu(menu)
	return rows, err
}

// 删除菜单
func (m MenuService) DropMenu(id string) error {
	err := NewMenuService().menuDao.DropMenu(id)
	return err
}

// 修改菜单
func (m MenuService) UpdateMenu(menu map[string]interface{}) error {
	err := NewMenuService().menuDao.UpdateMenu(menu)
	return err
}

// 获取菜单列表
func (m MenuService) GetMenuLists(menu entity.Menu) ([]entity.MenuMeta, error) {
	var metas []entity.MenuMeta
	lists, err := NewMenuService().menuDao.GetMenuLists(menu)
	//将菜单转成前端需要的格式
	for _, item := range lists {
		var meta = entity.MenuMeta{}
		meta.Meta.Id = item.ID
		meta.Meta.Title = item.Title
		meta.Meta.Icon = item.Icon
		meta.Meta.ParentId = item.ParentId
		meta.Meta.MenuType = item.MenuType
		meta.Meta.IsFrame = item.IsFrame
		meta.Meta.IsCache = item.IsCache
		meta.Meta.Visible = item.Visible
		meta.Meta.Status = item.Status
		meta.Meta.Perms = item.Perms
		meta.Meta.OrderNum = item.OrderNum
		meta.Meta.Query = item.Query
		meta.Meta.IsFull = item.IsFull
		meta.Meta.Remark = item.Remark
		meta.Meta.IsLink = item.IsLink
		meta.Path = item.Path
		meta.Name = item.Name
		meta.Component = item.Component
		meta.Redirect = item.Redirect
		metas = append(metas, meta)
	}
	//将菜单列表转成树形结构
	tree := utils.MenuTree(metas, "")

	return tree, err
}

// 获取添加菜单时需要选择的菜单树
func (m MenuService) GetMenuAddTree() ([]entity.AddMenuTree, error) {
	var menu = entity.Menu{}
	var menuTree []entity.AddMenuTree
	lists, err := NewMenuService().menuDao.GetMenuLists(menu)
	if err != nil {
		return nil, err
	}
	//将需要生成的树的结构体转成最终需要的结构体
	for _, item := range lists {
		var meta = entity.AddMenuTree{}
		meta.Id = item.ID
		meta.ParentId = item.ParentId
		meta.Title = item.Title
		meta.Name = item.Name
		menuTree = append(menuTree, meta)
	}
	//将菜单列表转成树形结构
	tree := utils.AddMenuTree(menuTree, "")
	return tree, err
}

// 根据id获取菜单
func (m MenuService) GetMenuById(id string) (*entity.Menu, error) {
	menu, err := NewMenuService().menuDao.GetMenuById(id)
	return menu, err
}
