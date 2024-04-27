package entity

type ColumnInfo struct {
	TABLE_SCHEMA             string `json:"tableSchema"`
	TABLE_NAME               string `json:"tableName"`
	COLUMN_NAME              string `json:"columnName"`
	COLUMN_COMMENT           string `json:"columnComment"`
	DATA_TYPE                string `json:"dataType"`
	COLUMN_TYPE              string `json:"columnType"`
	CHARACTER_MAXIMUM_LENGTH string `json:"characterMaximumLength"`
	CHARACTER_OCTET_LENGTH   string `json:"characterOctetLength"`
	COLUMN_KEY               string `json:"columnKey"`
	GoType                   string `json:"goType"`
	GoProperty               string `json:"goProperty"`
	QueryMethod              string `json:"queryMethod"`
	IsRequired               bool   `json:"isRequired"`
	HtmlType                 string `json:"htmlType"`
	DictType                 string `json:"dictType"`
}
