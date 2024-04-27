package dao

import (
	"go_study/src/common"
	"go_study/src/sys/dict/entity"
)

var dictTypeDao *DictTypeDao

type DictTypeDao struct {
	common.BaseDao
}

func NewDictTypeDao() *DictTypeDao {
	if dictTypeDao == nil {
		return &DictTypeDao{
			common.NewBaseDao(),
		}
	}
	return dictTypeDao
}

// 获取所有字典类型
func (DictTypeDao) GetAllDictType() ([]entity.DictType, error) {
	var dictType []entity.DictType
	tx := NewDictTypeDao().Orm.Find(&dictType)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return dictType, nil
}

// 根据数据字典的类型查询字典值
func (DictTypeDao) GetDictDataByDictType(dictType string) ([]entity.DictData, error) {
	var dictData []entity.DictData
	tx := NewDictTypeDao().Orm.Where("dict_type=?", dictType).Order("dict_sort").Find(&dictData)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return dictData, nil
}

// 获取字典类型树
func (DictTypeDao) GetDictTypeTree() ([]entity.DictType, error) {
	var dictType []entity.DictType
	tx := NewDictTypeDao().Orm.Find(&dictType)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return dictType, nil
}

// 根据Id查询字典类型信息
func (DictTypeDao) GetDictTypeById(dictId string) (entity.DictType, error) {
	var dictType entity.DictType
	tx := NewDictTypeDao().Orm.Where("id=?", dictId).Find(&dictType)
	if err := tx.Error; err != nil {
		return dictType, err
	}
	return dictType, nil
}

// 新增字典类型
func (DictTypeDao) AddDictType(dictType entity.DictType) error {
	tx := NewDictTypeDao().Orm.Create(&dictType)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

// 修改字典类型
func (DictTypeDao) UpdateDictType(dictType entity.DictType) error {
	tx := NewDictTypeDao().Orm.Save(&dictType)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

// 根据id切片删除字典类型
func (DictTypeDao) DropDictTypeByIds(ids []string) error {
	tx := NewDictTypeDao().Orm.Delete(&entity.DictType{}, ids)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

// 修改字典值
func (DictTypeDao) UpdateDictValueData(dictData entity.DictData) error {
	tx := NewDictTypeDao().Orm.Save(&dictData)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

// 新增字典值
func (DictTypeDao) AddDictValueData(dictData entity.DictData) error {
	tx := NewDictTypeDao().Orm.Create(&dictData)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

// 根据 ID 切片删除字典值
func (DictTypeDao) DropDictValueDataByIds(ids []string) (int64, error) {
	tx := NewDictTypeDao().Orm.Delete(&entity.DictData{}, ids)
	if err := tx.Error; err != nil {
		return 0, err
	}
	return tx.RowsAffected, nil
}
