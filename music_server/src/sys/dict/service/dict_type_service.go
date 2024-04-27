package service

import (
	"go_study/src/common"
	"go_study/src/sys/dict/dao"
	"go_study/src/sys/dict/entity"
	"go_study/src/utils"
)

var dictTypeService *DictTypeService

type DictTypeService struct {
	common.BaseService
	dictTypeDao *dao.DictTypeDao
}

func NewDictTypeService() *DictTypeService {
	if dictTypeService == nil {
		return &DictTypeService{
			dictTypeDao: dao.NewDictTypeDao(),
		}
	}
	return dictTypeService
}

// 获取所有字典类型
func (DictTypeService) GetAllDictType() ([]entity.DictType, error) {
	dictType, err := NewDictTypeService().dictTypeDao.GetAllDictType()
	return dictType, err
}

// 根据数据字典的类型查询字典值
func (DictTypeService) GetDictDataByDictType(dictType string) ([]entity.DictData, error) {
	data, err := NewDictTypeService().dictTypeDao.GetDictDataByDictType(dictType)
	return data, err
}

// 获取字典类型树
func (DictTypeService) GetDictTypeTree() ([]entity.DictType, error) {
	data, err := NewDictTypeService().dictTypeDao.GetDictTypeTree()
	if err != nil {
		return nil, err
	}
	//将字典类型列表转化成字典类型树
	tree := utils.DictTypeTree(data, "0")
	return tree, err
}

// 根据Id查询字典类型信息
func (DictTypeService) GetDictTypeById(dictId string) (entity.DictType, error) {
	data, err := NewDictTypeService().dictTypeDao.GetDictTypeById(dictId)
	return data, err
}

// 新增字典类型
func (DictTypeService) AddDictType(dictType entity.DictType) error {
	dictType.ID = utils.GenerateID()
	err := NewDictTypeService().dictTypeDao.AddDictType(dictType)
	return err
}

// 修改字典类型
func (DictTypeService) UpdateDictType(dictType entity.DictType) error {
	err := NewDictTypeService().dictTypeDao.UpdateDictType(dictType)
	return err
}

// 根据id切片删除字典类型
func (DictTypeService) DropDictTypeByIds(ids []string) error {
	err := NewDictTypeService().dictTypeDao.DropDictTypeByIds(ids)
	return err
}

// 更新字典值
func (DictTypeService) UpdateDictValueData(dictData entity.DictData) error {
	err := NewDictTypeService().dictTypeDao.UpdateDictValueData(dictData)
	return err
}

// 新增字典值
func (DictTypeService) AddDictValueData(dictData entity.DictData) error {
	//生成id
	dictData.ID = utils.GenerateID()
	err := NewDictTypeService().dictTypeDao.AddDictValueData(dictData)
	return err
}

// 根据id切片删除字典值
func (DictTypeService) DropDictValueDataByIds(ids []string) (int64, error) {
	rows, err := NewDictTypeService().dictTypeDao.DropDictValueDataByIds(ids)
	return rows, err
}
