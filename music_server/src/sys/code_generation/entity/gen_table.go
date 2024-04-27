package entity

import "go_study/src/common"

type GenTable struct {
	common.Model
	TableSchema    string `gorm:"type:varchar(64);not null;comment:表所属数据库" json:"tableSchema"`
	TablesName     string `gorm:"type:varchar(200);not null;comment:表名称" json:"tableName"`
	TableComment   string `gorm:"type:varchar(500);comment:表描述" json:"tableComment"`
	SubTableName   string `gorm:"type:varchar(200);comment:子表关联名称" json:"subTableName"`
	SubTableFkName string `gorm:"type:varchar(64);comment:子表关联的外键名" json:"subTableFkName"`
	StructName     string `gorm:"type:varchar(100);comment:结构体名称" json:"structName"`
	TemplateUse    string `gorm:"type:varchar(10);comment:使用的模板（crud单表操作 tree树表操作）" json:"templateUse"`
	PackageName    string `gorm:"type:varchar(100);comment:生成包名称" json:"packageName"`
	AuthorName     string `gorm:"type:varchar(100);comment:生成作者名称" json:"authorName"`
	GenType        string `gorm:"type:char(1);comment:生成代码方式（0zip压缩包 1自定义路径）" json:"genType"`
	GenPath        string `gorm:"type:varchar(200);comment:生成路径（不填则默认在当前项目下生成）" json:"genPath"`
	Options        string `gorm:"type:text;comment:其它生成选项" json:"options"`
	Remark         string `gorm:"type:varchar(500);comment:备注" json:"remark"`
}

func (GenTable) TableName() string {
	return "tb_gen_table"
}
