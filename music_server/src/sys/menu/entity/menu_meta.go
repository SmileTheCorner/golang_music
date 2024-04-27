package entity

type MenuMeta struct {
	Path      string     `json:"path"`
	Name      string     `json:"name"`
	Component string     `json:"component"`
	Redirect  string     `json:"redirect"`
	Meta      MetaProps  `json:"meta"`
	Children  []MenuMeta `json:"children"`
}

type MetaProps struct {
	Id       string `json:"id"`
	ParentId string `json:"parentId"`
	Icon     string `json:"icon"`
	Title    string `json:"title"`
	IsLink   string `json:"isLink"`
	Visible  int8   `json:"visible"`
	IsFull   int8   `json:"isFull"`
	IsFrame  int8   `json:"isFrame"`
	IsCache  int8   `json:"isCache"`
	Query    string `json:"query"`
	Perms    string `json:"perms"`
	Status   int8   `json:"status"`
	OrderNum int8   `json:"orderNum"`
	Remark   string `json:"remark"`
	MenuType string `json:"menuType"`
}
