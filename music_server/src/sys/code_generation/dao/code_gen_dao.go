package dao

import (
	"go_study/src/common"
	"go_study/src/sys/code_generation/entity"
	"go_study/src/utils"
	"strings"
)

type CodeGenDao struct {
	common.BaseDao
}

var codeGenDao *CodeGenDao

// NewCodeGenDao 做成单实列
func NewCodeGenDao() *CodeGenDao {
	if codeGenDao == nil {
		return &CodeGenDao{
			common.NewBaseDao(),
		}
	}
	return codeGenDao
}

// 查询当前连接数据库的所有表命名和表名注释
func (m CodeGenDao) GetCurrentAllTables(page common.Page) ([]entity.TableInfo, int64, error) {
	var tables []entity.TableInfo
	var count int64
	queryMap := utils.PageToObject(&page)
	var builder strings.Builder
	for key, value := range queryMap {
		builder.WriteString(" and ")
		keyUpper := strings.ToUpper(key)
		builder.WriteString(keyUpper)
		builder.WriteString(" = ")
		builder.WriteString("'" + value.(string) + "'")
		builder.WriteString(" ")
	}
	tx := NewCodeGenDao().Orm.Raw("select TABLE_SCHEMA,TABLE_NAME,CREATE_TIME,UPDATE_TIME,TABLE_COMMENT from information_schema.TABLES where TABLE_SCHEMA=(select database())"+builder.String()+" LIMIT ?,?", page.Offset, page.PageSize).Scan(&tables)
	return tables, count, tx.Error
}

// 查询当前连接数据库的所有列的名称、类型、注释等
func (m CodeGenDao) GetCurrentAllColumns(tableName string) ([]entity.ColumnInfo, error) {
	var columns []entity.ColumnInfo
	tx := NewCodeGenDao().Orm.Raw("select TABLE_SCHEMA,TABLE_NAME,COLUMN_NAME,COLUMN_COMMENT,DATA_TYPE,COLUMN_TYPE,CHARACTER_MAXIMUM_LENGTH,CHARACTER_OCTET_LENGTH,COLUMN_KEY from information_schema.COLUMNS where TABLE_SCHEMA=(select database()) and TABLE_NAME = ?", tableName).Scan(&columns)
	return columns, tx.Error
}

// 根据表Ids切片查询表信息
func (m CodeGenDao) GetGenTableByIds(ids []string) (entity.GenCodeTableInfo, error) {
	//查询表信息
	var genTables []entity.GenTable
	var columns []entity.GenTableColumn
	tableTx := NewCodeGenDao().Orm.Where("id in (?)", ids).Find(&genTables)
	//查询列信息
	columnsTx := NewCodeGenDao().Orm.Where("table_id in (?)", ids).Find(&columns)

	if tableTx.Error != nil {
		return entity.GenCodeTableInfo{}, tableTx.Error
	}
	if columnsTx.Error != nil {
		return entity.GenCodeTableInfo{}, columnsTx.Error
	}
	var data = entity.GenCodeTableInfo{
		TableInfo:  genTables,
		ColumnInfo: columns,
	}
	return data, nil
}

// 导入代码生成业务表
func (m CodeGenDao) ImportGenTable(genTable []entity.GenTable) (int64, error) {
	//生成ID
	for i, v := range genTable {
		id := utils.GenerateID()
		split := strings.Split(v.TablesName, "_")
		var builder strings.Builder
		for _, value := range split {
			if !strings.Contains(strings.ToLower(value), strings.ToLower("tb")) {
				builder.WriteString(utils.UpperCaseFirst(value))
			}

		}
		genTable[i].ID = id
		genTable[i].StructName = builder.String()
	}
	tx := NewCodeGenDao().Orm.Create(&genTable)
	return tx.RowsAffected, tx.Error
}

// 分页查询代码生成业务表
func (m CodeGenDao) QueryGenTableList(page common.Page) ([]entity.GenTable, int64, error) {
	var genTable []entity.GenTable
	var count int64
	queryMap := utils.PageToObject(&page)
	tx := NewCodeGenDao().Orm.Where(queryMap).Limit(page.PageSize).Offset(page.Offset).Find(&genTable).Offset(-1).Limit(-1).Count(&count)
	return genTable, count, tx.Error
}

// 根据ID切片删除代码生成业务表
func (m CodeGenDao) DeleteGenTableByIds(ids []string) (int64, error) {
	tx := NewCodeGenDao().Orm.Delete(&entity.GenTable{}, ids)
	return tx.RowsAffected, tx.Error
}
