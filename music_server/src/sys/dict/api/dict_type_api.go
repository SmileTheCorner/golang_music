package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go_study/src/sys/dict/entity"
	"go_study/src/sys/dict/service"
	"go_study/src/utils"
)

type DictTypeApi struct {
	dictTypeService *service.DictTypeService
}

func NewDictTypeApi() *DictTypeApi {
	return &DictTypeApi{
		dictTypeService: service.NewDictTypeService(),
	}
}

// 获取所有字典类型
func (DictTypeApi) GetAllDictType(ctx *gin.Context) {
	dictType, err := NewDictTypeApi().dictTypeService.GetAllDictType()
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, dictType)
	return
}

// 根据数据字典的类型查询字典值
func (DictTypeApi) GetDictDataByDictType(ctx *gin.Context) {
	dictType, b := ctx.GetQuery("dictType")
	if b {
		byDictType, err := NewDictTypeApi().dictTypeService.GetDictDataByDictType(dictType)
		if err != nil {
			utils.Fail(ctx, err)
			return
		}
		utils.Ok(ctx, byDictType)
		return
	}
	utils.Fail(ctx, errors.New("参数传递错误"))
}

// 获取字典类型树
func (DictTypeApi) GetDictTypeTree(ctx *gin.Context) {
	data, err := NewDictTypeApi().dictTypeService.GetDictTypeTree()
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, data)
	return
}

// 根据Id查询字典类型信息
func (DictTypeApi) GetDictTypeById(ctx *gin.Context) {
	dictId, b := ctx.GetQuery("id")
	if b {
		data, err := NewDictTypeApi().dictTypeService.GetDictTypeById(dictId)
		if err != nil {
			utils.Fail(ctx, err)
			return
		}
		utils.Ok(ctx, data)
		return
	}
	utils.Fail(ctx, errors.New("参数传递错误"))
}

// 新增字典类型
func (DictTypeApi) AddDictType(ctx *gin.Context) {
	var dictType entity.DictType
	err := ctx.ShouldBind(&dictType)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	err = NewDictTypeApi().dictTypeService.AddDictType(dictType)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, nil)
	return
}

// 修改字典类型
func (DictTypeApi) UpdateDictType(ctx *gin.Context) {
	var dictType entity.DictType
	err := ctx.ShouldBind(&dictType)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	err = NewDictTypeApi().dictTypeService.UpdateDictType(dictType)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, nil)
	return
}

// 根据id切片删除字典类型
func (DictTypeApi) DropDictTypeByIds(ctx *gin.Context) {
	var ids []string
	err := ctx.ShouldBind(&ids)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	err = NewDictTypeApi().dictTypeService.DropDictTypeByIds(ids)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, nil)
	return
}

// 更改字典值
func (DictTypeApi) UpdateDictValueData(ctx *gin.Context) {
	var dictData entity.DictData
	err := ctx.ShouldBind(&dictData)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	err = NewDictTypeApi().dictTypeService.UpdateDictValueData(dictData)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, nil)
	return
}

// 新增字典值
func (DictTypeApi) AddDictValueData(ctx *gin.Context) {
	var dictData entity.DictData
	err := ctx.ShouldBind(&dictData)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	err = NewDictTypeApi().dictTypeService.AddDictValueData(dictData)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, nil)
	return
}

// 根据 id 切片删除字典值
func (DictTypeApi) DropDictValueDataByIds(ctx *gin.Context) {
	var ids []string
	err := ctx.ShouldBind(&ids)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	rows, err := NewDictTypeApi().dictTypeService.DropDictValueDataByIds(ids)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, rows)
	return
}
