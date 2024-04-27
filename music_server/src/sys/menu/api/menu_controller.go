package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go_study/src/sys/menu/entity"
	"go_study/src/sys/menu/service"
	"go_study/src/utils"
)

type MenuApi struct {
	menuService *service.MenuService
}

func NewMenuApi() *MenuApi {
	return &MenuApi{
		menuService: service.NewMenuService(),
	}
}

// 添加菜单
func (m MenuApi) AddMenu(ctx *gin.Context) {
	var menu entity.Menu
	//获取添加的菜单参数
	err := ctx.ShouldBind(&menu)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	addMenu, err := NewMenuApi().menuService.AddMenu(menu)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, addMenu)
}

// 删除菜单
func (m MenuApi) DropMenu(ctx *gin.Context) {
	//获取要删除菜单的id
	id, b := ctx.GetQuery("id")
	if !b {
		utils.Fail(ctx, errors.New("参数传递错误"))
		return
	}
	err := NewMenuApi().menuService.DropMenu(id)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, "删除成功")
}

// 修改菜单
func (m MenuApi) UpdateMenu(ctx *gin.Context) {
	menu := make(map[string]interface{})
	err := ctx.BindJSON(&menu)
	menu = utils.MapKeyConvUnderline(menu)
	if err != nil {
		utils.Fail(ctx, errors.New("参数绑定错误"))
		return
	}
	err = NewMenuApi().menuService.UpdateMenu(menu)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, "修改成功")
}

// 获取菜单列表
func (m MenuApi) GetMenuLists(ctx *gin.Context) {
	var menu entity.Menu
	//获取参数
	ctx.ShouldBind(&menu)
	lists, err := NewMenuApi().menuService.GetMenuLists(menu)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, lists)
}

// 获取添加菜单时的选择的菜单树
func (m MenuApi) GetMenuAddTree(ctx *gin.Context) {
	tree, err := NewMenuApi().menuService.GetMenuAddTree()
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, tree)
}

// 根据ID获取菜单
func (m MenuApi) GetMenuById(ctx *gin.Context) {
	//从请求路径中获取id
	idStr, b := ctx.GetQuery("id")
	if b {
		menu, err := NewMenuApi().menuService.GetMenuById(idStr)
		if err != nil {
			utils.Fail(ctx, err)
			return
		} else {
			utils.Ok(ctx, menu)
		}
	} else {
		utils.Fail(ctx, errors.New("参数中没有id"))
		return
	}
}
