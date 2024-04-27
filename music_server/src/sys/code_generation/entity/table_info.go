package entity

type TableInfo struct {
	TABLE_SCHEMA  string `json:"tableSchema"`
	TABLE_NAME    string `json:"tableName"`
	CREATE_TIME   string `json:"createTime"`
	UPDATE_TIME   string `json:"updateTime"`
	TABLE_COMMENT string `json:"tableComment"`
}
