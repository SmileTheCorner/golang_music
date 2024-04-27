package entity

import "go_study/src/common"

type GenTableColumn struct {
	common.Model
	TableId       string `gorm:"type:varchar(200);not null;comment:归属表编号" json:"tableId"`
	ColumnName    string `gorm:"type:varchar(200);not null;comment:列名称" json:"columnName"`
	ColumnComment string `gorm:"type:varchar(500);comment:列描述" json:"columnComment"`
	ColumnType    string `gorm:"type:varchar(100);comment:列类型" json:"columnType"`
	GoType        string `gorm:"type:varchar(100);comment:Go类型" json:"goType"`
	GoField       string `gorm:"type:varchar(100);comment:Go字段名" json:"goField"`
	IsPk          string `gorm:"type:char(1);comment:是否主键字段（1是 0否）" json:"isPk"`
	IsIncrement   string `gorm:"type:char(1);comment:是否自增字段（1是 0否）" json:"isIncrement"`
	IsRequired    string `gorm:"type:char(1);comment:是否必填字段（1是 0否）" json:"isRequired"`
	IsQuery       string `gorm:"type:char(1);comment:是否查询字段（1是 0否）" json:"isQuery"`
	QueryType     string `gorm:"type:char(1);comment:查询方式（0等于、1不等于、2大于、3小于、4范围、5模糊）" json:"queryType"`
	HtmlType      string `gorm:"type:varchar(50);comment:显示类型（0密码框 1单行文本 2多行文本 3富文本 4单选框 5复选框 6下拉框 7日期控件 8时间控件 9日期时间控件）" json:"htmlType"`
	DictType      string `gorm:"type:varchar(200);comment:字典类型" json:"dictType"`
	Sort          int    `gorm:"type:int;comment:排序" json:"sort"`
}

func (GenTableColumn) TableName() string {
	return "tb_gen_table_column"
}
