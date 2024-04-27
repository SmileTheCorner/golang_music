package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go_study/src/common"
	"go_study/src/sys/user/entity"
	"go_study/src/sys/user/service"
	"go_study/src/utils"
)

type userApi struct {
	userService *service.UserService
}

func NewUserApi() userApi {
	return userApi{
		userService: service.NewUserService(),
	}
}

// Login
// @Tag 用户管理模块
// @Summary 用户登录
// @Description 用户名和密码登录
// @Param name formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {string} string "登录成功"
// @Failure 400 {string} string "登录失败"
// @Router /api/v1/public/user/login [post]
func (m userApi) Login(ctx *gin.Context) {
	var paramsMap = make(map[string]string)
	ctx.ShouldBindJSON(&paramsMap)
	userName := paramsMap["username"]
	password := paramsMap["password"]
	token, err := NewUserApi().userService.Login(userName, password)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	var userMap map[string]string = map[string]string{"access_token": token}
	utils.Ok(ctx, userMap)
}

// AddUser 添加用户
// @Summary 新增用户
// @Description  新增用户信息
// @Param Name formData string  true "用户名"
// @Param RealName formData string false "真实姓名"
// @Param Avatar formData string false "头像"
// @Param Mobile formData string false "电话"
// @Param Email formData string false "邮箱"
// @Param Password formData string false "密码"
// @Success 200 {object} utils.ResponseResult
// @Router /api/v1/user/addUser [post]
func (userApi) AddUser(ctx *gin.Context) {
	var user entity.User
	ctx.ShouldBind(&user)
	isUser := NewUserApi().userService.QueryUserByUserName(user.Name)
	if isUser.Name != "" {
		utils.Fail(ctx, errors.New("该用户名已经存在"))
		return
	}
	//调用service层的添加用户方法
	err := NewUserApi().userService.AddUserService(user)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	utils.Ok(ctx, nil)
}

// QueryUserByUserName 根据用户名查找用户
// @Summary 根据用户名查找用户
// @Description  根据用户名查询用户信息
// @Param userName query string  true "用户名"
// @Success 200 {object} utils.ResponseResult
// @Router /api/v1/user/queryUserByName [get]
func (userApi) QueryUserByUserName(ctx *gin.Context) {
	userName := ctx.Query("userName")
	user := NewUserApi().userService.QueryUserByUserName(userName)
	if user.Name == "" {
		utils.Fail(ctx, errors.New("该用户不存在"))
		return
	}
	utils.Ok(ctx, user)
}

// QueryUserList 分页查询用户列表
func (userApi) QueryUserList(ctx *gin.Context) {
	var page common.Page
	ctx.ShouldBind(&page)
	list, count, err := NewUserApi().userService.QueryUserList(page)
	if err != nil {
		utils.Fail(ctx, err)
		return
	}
	data := utils.PageData(list, int64(page.PageNum), int64(page.PageSize), count)
	utils.Ok(ctx, data)
}
