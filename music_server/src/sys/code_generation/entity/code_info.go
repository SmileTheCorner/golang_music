package entity

type CodeInfo struct {
	ModuleTitle string           `json:"moduleTitle"`
	Author      string           `json:"author"`
	Date        string           `json:"date"`
	ModuleName  string           `json:"moduleName"`
	Fields      []GenTableColumn `json:"fields"`
}
