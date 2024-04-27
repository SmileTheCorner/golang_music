package service

import (
	"errors"
	"go_study/src/common"
	"go_study/src/sys/user/dao"
	"go_study/src/sys/user/entity"
	"go_study/src/utils"
)

var userService *UserService

type UserService struct {
	UserDao *dao.UserDao
	common.BaseService
}

func NewUserService() *UserService {
	if userService == nil {
		return &UserService{
			UserDao: dao.NewUserDao(),
		}
	}
	return userService
}

// AddUserService 添加用户
func (m UserService) AddUserService(user entity.User) error {
	//对密码加密
	str := utils.SM4Encrypt(user.Password)
	user.Password = str
	err := NewUserService().UserDao.AddUser(user)
	return err
}

// QueryUserByUserName 根据用户名查找用户
func (m UserService) QueryUserByUserName(userName string) entity.User {
	return NewUserService().UserDao.QueryUserByUserName(userName)
}

// Login 用户登录
func (m UserService) Login(userName, password string) (string, error) {
	//根据用户名去查找用户信息
	user := NewUserService().UserDao.QueryUserByUserName(userName)
	if user.Name == "" {
		return "", errors.New("该用户不存在")
	}
	//对密码进行加密
	pass := utils.SM4Encrypt(password)
	if pass == user.Password {
		//用户名和密码正确则生成token
		token, err := utils.GenerateToken(user.ID, user.Name)
		return token, err
	} else {
		return "", errors.New("用户名或密码不存在")
	}

}

// QueryUserList 分页查询用户列表
func (m UserService) QueryUserList(page common.Page) ([]entity.User, int64, error) {
	return m.UserDao.QueryUserList(page)
}
