package utils

import (
	dep "go_study/src/sys/department/entity"
	dictEntity "go_study/src/sys/dict/entity"
	"go_study/src/sys/menu/entity"
)

// MenuTree 菜单树
func MenuTree(menus []entity.MenuMeta, parentId string) []entity.MenuMeta {
	res := make([]entity.MenuMeta, 0)
	for _, menu := range menus {
		//获取到父节点
		pid := menu.Meta.ParentId
		if pid == parentId {
			menu.Children = MenuTree(menus, menu.Meta.Id)
			res = append(res, menu)
		}
	}
	return res
}

// AddMenuTree 添加菜单树
func AddMenuTree(menus []entity.AddMenuTree, parentId string) []entity.AddMenuTree {
	res := make([]entity.AddMenuTree, 0)
	for _, menu := range menus {
		//获取到父节点
		pid := menu.ParentId
		if pid == parentId {
			menu.Children = AddMenuTree(menus, menu.Id)
			res = append(res, menu)
		}
	}
	return res
}

// 将部门列表转成菜单树
func DepartmentTree(departments []dep.Department, parentId string) []dep.Department {
	res := make([]dep.Department, 0)
	for _, department := range departments {
		//获取到父节点
		pid := department.ParentId
		if pid == parentId {
			department.Children = DepartmentTree(departments, department.ID)
			res = append(res, department)
		}
	}
	return res
}

// 将字典列表转化成字典类型树
func DictTypeTree(dictTypes []dictEntity.DictType, parentId string) []dictEntity.DictType {
	res := make([]dictEntity.DictType, 0)
	for _, dictType := range dictTypes {
		//获取到父节点
		pid := dictType.ParentId
		if pid == parentId {
			dictType.Children = DictTypeTree(dictTypes, dictType.ID)
			res = append(res, dictType)
		}
	}
	return res
}
