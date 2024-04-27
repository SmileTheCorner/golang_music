package service

import (
	"bytes"
	"go_study/src/common"
	"go_study/src/sys/code_generation/dao"
	"go_study/src/sys/code_generation/entity"
	"go_study/src/utils"
	"io"
	"text/template"
	"time"
)

type CodeGenService struct {
	common.BaseService
	codeGenDao *dao.CodeGenDao
}

var codeGenService *CodeGenService

// NewCodeGenService 做成单实列
func NewCodeGenService() *CodeGenService {
	if codeGenService == nil {
		return &CodeGenService{
			codeGenDao: dao.NewCodeGenDao(),
		}
	}
	return codeGenService
}

// 获取当前数据库所有表信息
func (m CodeGenService) GetCurrentAllTables(page common.Page) ([]entity.TableInfo, int64, error) {
	tables, count, err := NewCodeGenService().codeGenDao.GetCurrentAllTables(page)
	return tables, count, err
}

// 导入代码生成业务表
func (m CodeGenService) ImportGenTable(genTable []entity.GenTable) (int64, error) {
	data, err := NewCodeGenService().codeGenDao.ImportGenTable(genTable)
	return data, err
}

// 分页查询代码生成业务表
func (m CodeGenService) GetGenTableList(page common.Page) ([]entity.GenTable, int64, error) {
	data, count, err := NewCodeGenService().codeGenDao.QueryGenTableList(page)
	return data, count, err
}

// 根据表名查询表信息
func (m CodeGenService) GetCurrentAllColumns(tableName string) ([]entity.ColumnInfo, error) {
	data, err := NewCodeGenService().codeGenDao.GetCurrentAllColumns(tableName)
	//将mysql转成对应的go类型
	for i := 0; i < len(data); i++ {
		data[i].GoType = utils.MysqlTypeToGoType(data[i].DATA_TYPE)
		data[i].GoProperty = utils.RemoveUnderline(data[i].COLUMN_NAME)
		data[i].QueryMethod = "EQ"
		data[i].IsRequired = false
		data[i].HtmlType = "input"
		data[i].DictType = ""
	}
	return data, err
}

// 根据表Ids切片查询表信息
func (m CodeGenService) GetGenTableByIds(ids []string, buf io.Writer) error {
	data, err := NewCodeGenService().codeGenDao.GetGenTableByIds(ids)
	//根据查询到的data生成模版
	var codeInfo = entity.CodeInfo{
		ModuleTitle: data.TableInfo[0].TableComment,
		Author:      "admin",
		Date:        time.Now().Format("2006-01-02 15:04:05"),
		ModuleName:  data.TableInfo[0].StructName,
		Fields:      data.ColumnInfo,
	}
	code, err := m.GeneratorCode(codeInfo)
	if err != nil {
		return err
	}
	//获取到要压缩的内容
	err = utils.ZipFiles("code.zip", code, buf)
	if err != nil {
		return err
	}
	return nil
}

// 根据ID切片删除代码生成业务表
func (m CodeGenService) DeleteGenTableByIds(ids []string) (int64, error) {
	data, err := NewCodeGenService().codeGenDao.DeleteGenTableByIds(ids)
	return data, err
}

// 根据模版生成代码
func (m CodeGenService) GeneratorCode(data entity.CodeInfo) (map[string]string, error) {
	//返回的数据
	resultData := make(map[string]string)
	//创建自定义的模版
	daoTemplate := template.New("dao.tmpl").Funcs(template.FuncMap{
		"toLowerFirst": utils.LowerCaseFirst,
	})
	//解析dao模板
	daoTemplate, err := daoTemplate.ParseFiles("src/sys/code_generation/template/dao.tmpl")
	if err != nil {
		return nil, err
	}
	var daoCode bytes.Buffer
	err = daoTemplate.Execute(&daoCode, data)
	if err != nil {
		return nil, err
	}

	//创建自定义的模版
	apiTemplate := template.New("api.tmpl").Funcs(template.FuncMap{
		"toLowerFirst": utils.LowerCaseFirst,
	})
	//解析api模板
	apiTemplate, err = apiTemplate.ParseFiles("src/sys/code_generation/template/api.tmpl")
	if err != nil {
		return nil, err
	}
	var apiCode bytes.Buffer
	err = apiTemplate.Execute(&apiCode, data)
	if err != nil {
		return nil, err
	}
	//创建自定义的模版
	serviceTemplate := template.New("service.tmpl").Funcs(template.FuncMap{
		"toLowerFirst": utils.LowerCaseFirst,
	})
	//解析service模板
	serviceTemplate, err = serviceTemplate.ParseFiles("src/sys/code_generation/template/service.tmpl")
	if err != nil {
		return nil, err
	}
	var serviceCode bytes.Buffer
	err = serviceTemplate.Execute(&serviceCode, data)
	if err != nil {
		return nil, err
	}
	//创建自定义的模版
	entityTemplate := template.New("entity.tmpl").Funcs(template.FuncMap{
		"toLowerFirst":      utils.LowerCaseFirst,
		"toUpperFirst":      utils.UpperCaseFirst,
		"mysqlTypeToGoType": utils.MysqlTypeToGoType,
		"removeUnderline":   utils.RemoveUnderline,
	})
	//解析entity模板
	entityTemplate, err = entityTemplate.ParseFiles("src/sys/code_generation/template/entity.tmpl")
	if err != nil {
		return nil, err
	}
	var entityCode bytes.Buffer
	err = entityTemplate.Execute(&entityCode, data)
	if err != nil {
		return nil, err
	}

	//创建自定义的模版
	routerTemplate := template.New("router.tmpl").Funcs(template.FuncMap{
		"toLowerFirst": utils.LowerCaseFirst,
	})
	//解析router模板
	routerTemplate, err = routerTemplate.ParseFiles("src/sys/code_generation/template/router.tmpl")
	if err != nil {
		return nil, err
	}
	var routerCode bytes.Buffer
	err = routerTemplate.Execute(&routerCode, data)
	if err != nil {
		return nil, err
	}

	//创建自定义的模版
	axiosTemplate := template.New("axios.tmpl").Funcs(template.FuncMap{
		"toLowerFirst": utils.LowerCaseFirst,
	})
	//解析axios模板
	axiosTemplate, err = axiosTemplate.ParseFiles("src/sys/code_generation/template/axios.tmpl")
	if err != nil {
		return nil, err
	}
	var axiosCode bytes.Buffer
	err = axiosTemplate.Execute(&axiosCode, data)
	if err != nil {
		return nil, err
	}

	//创建自定义的模版
	indexTemplate := template.New("index.tmpl").Funcs(template.FuncMap{
		"toLowerFirst": utils.LowerCaseFirst,
	})
	//解析index模板
	indexTemplate, err = indexTemplate.ParseFiles("src/sys/code_generation/template/index.tmpl")
	if err != nil {
		return nil, err
	}
	var indexCode bytes.Buffer
	err = indexTemplate.Execute(&indexCode, data)
	if err != nil {
		return nil, err
	}

	//创建自定义的模版
	typeTemplate := template.New("type.tmpl").Funcs(template.FuncMap{
		"removeUnderline":   utils.RemoveUnderline,
		"mysqlTypeToTsType": utils.MysqlTypeToTsType,
	})
	//解析type模板
	typeTemplate, err = typeTemplate.ParseFiles("src/sys/code_generation/template/type.tmpl")
	if err != nil {
		return nil, err
	}
	var typeCode bytes.Buffer
	err = typeTemplate.Execute(&typeCode, data)
	if err != nil {
		return nil, err
	}

	resultData["daoCode"] = daoCode.String()
	resultData["serviceCode"] = serviceCode.String()
	resultData["controllerCode"] = apiCode.String()
	resultData["entityCode"] = entityCode.String()
	resultData["routerCode"] = routerCode.String()
	resultData["axiosCode"] = axiosCode.String()
	resultData["indexCode"] = indexCode.String()
	resultData["typeCode"] = typeCode.String()
	return resultData, nil
}
