package entity

// 添加菜单时需要选择的菜单树
type AddMenuTree struct {
	Id       string        `json:"value"`
	ParentId string        `json:"parentId"`
	Title    string        `json:"label"`
	Name     string        `json:"name"`
	Children []AddMenuTree `json:"children"`
}
