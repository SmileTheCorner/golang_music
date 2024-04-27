package dao

import (
	"errors"
	"go_study/src/common"
	"go_study/src/global"
	"go_study/src/sys/user/entity"
	"go_study/src/utils"
)

var userDao *UserDao

type UserDao struct {
	common.BaseDao
}

// NewUserDao 初始化UserDao
func NewUserDao() *UserDao {
	if userDao == nil {
		return &UserDao{
			common.NewBaseDao(),
		}
	}
	return userDao
}

// AddUser 新增用户
func (m UserDao) AddUser(user entity.User) error {
	err := global.DB.Create(&user).Error
	if err != nil {
		return errors.New("add user error")
	}
	return nil
}

// QueryUserByUserName 根据用户名查询用户
func (m UserDao) QueryUserByUserName(userName string) entity.User {
	var user entity.User
	m.Orm.Model(&user).Where("name=?", userName).Select("id,name,password,real_name, avatar, mobile, email").Find(&user)
	return user
}

// 分页查询用户列表
func (m UserDao) QueryUserList(page common.Page) ([]entity.User, int64, error) {
	var count int64
	var user []entity.User
	queryMap := utils.PageToObject(&page)
	tx := m.Orm.Where(queryMap).Offset(page.Offset).Limit(page.PageSize).Find(&user).Limit(-1).Offset(-1).Count(&count)
	return user, count, tx.Error

}
