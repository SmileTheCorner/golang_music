package entity

type GenCodeTableInfo struct {
	TableInfo  []GenTable       `json:"tableInfo"`
	ColumnInfo []GenTableColumn `json:"columnInfo"`
}
