package dao

import (
	"go_study/src/common"
	"go_study/src/sys/department/entity"
	"go_study/src/utils"
)

type DepartmentDao struct {
	common.BaseDao
}

var departmentDao *DepartmentDao

// NewDepartmentDao 做成单实列
func NewDepartmentDao() *DepartmentDao {
	if departmentDao == nil {
		return &DepartmentDao{
			common.NewBaseDao(),
		}
	}
	return departmentDao
}

// 获取部门列表
func (m DepartmentDao) GetDepartmentLists(department entity.Department, page common.Page) ([]entity.Department, int64, error) {
	var lists []entity.Department
	var count int64
	offset := (page.PageNum - 1) * page.PageSize
	tx := NewDepartmentDao().Orm.Where(department).Offset(offset).Limit(page.PageSize).Find(&lists).Offset(-1).Limit(-1).Count(&count)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	return lists, count, nil
}

// 添加部门信息
func (m DepartmentDao) AddDepartment(department entity.Department) error {
	//生成id
	department.ID = utils.GenerateID()
	tx := NewDepartmentDao().Orm.Create(&department)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// 修改部门信息
func (m DepartmentDao) UpdateDepartment(department entity.Department) error {
	tx := NewDepartmentDao().Orm.Model(&department).Updates(department)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// 删除部门信息
func (m DepartmentDao) DropDepartment(id []string) error {
	var department entity.Department
	tx := NewDepartmentDao().Orm.Delete(&department, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// 查询部门树
func (m DepartmentDao) GetDepartmentTree() ([]entity.Department, error) {
	var lists []entity.Department
	tx := NewDepartmentDao().Orm.Find(&lists)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return lists, nil
}
