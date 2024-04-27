package service

import (
	"go_study/src/common"
	"go_study/src/sys/department/dao"
	"go_study/src/sys/department/entity"
	"go_study/src/utils"
)

type DepartmentService struct {
	departmentDao *dao.DepartmentDao
	common.BaseService
}

var departmentService *DepartmentService

// NewDepartmentService 做成单实列
func NewDepartmentService() *DepartmentService {
	if departmentService == nil {
		return &DepartmentService{
			departmentDao: dao.NewDepartmentDao(),
		}
	}
	return departmentService
}

// 获取部门列表
func (m DepartmentService) GetDepartmentLists(department entity.Department, page common.Page) ([]entity.Department, int64, error) {
	departmentLists, count, err := NewDepartmentService().departmentDao.GetDepartmentLists(department, page)
	if err != nil {
		return nil, 0, err
	}
	//将部门列表转换成树形结构
	tree := utils.DepartmentTree(departmentLists, "")
	return tree, count, nil
}

// 添加部门信息
func (m DepartmentService) AddDepartment(department entity.Department) error {
	err := NewDepartmentService().departmentDao.AddDepartment(department)
	if err != nil {
		return err
	}
	return nil
}

// 修改部门信息
func (m DepartmentService) UpdateDepartment(department entity.Department) error {
	err := NewDepartmentService().departmentDao.UpdateDepartment(department)
	if err != nil {
		return err
	}
	return nil
}

// 删除部门信息
func (m DepartmentService) DropDepartment(id []string) error {
	err := NewDepartmentService().departmentDao.DropDepartment(id)
	if err != nil {
		return err
	}
	return nil
}

// 获取部门树形结构
func (m DepartmentService) GetDepartmentTree() ([]entity.Department, error) {
	departmentLists, err := NewDepartmentService().departmentDao.GetDepartmentTree()
	if err != nil {
		return nil, err
	}
	//将部门列表转换成树形结构
	tree := utils.DepartmentTree(departmentLists, "")
	return tree, nil
}
