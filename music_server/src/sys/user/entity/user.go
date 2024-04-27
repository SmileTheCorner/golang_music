package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"size:64;not null;comment:用户名" json:"name"`
	RealName string `gorm:"size:128;comment:真实姓名" json:"realName"`
	Avatar   string `gorm:"size:255;comment:头像" json:"avatar"`
	Mobile   string `gorm:"size:11;comment:电话" json:"mobile"`
	Email    string `gorm:"size:128;comment:邮箱" json:"email"`
	Password string `gorm:"size:128;not null;comment:密码" json:"password"`
	Gender   int8   `gorm:"size:1;comment:性别" json:"gender"`
	IdCard   string `gorm:"size:20;comment:身份证号" json:"idCard"`
	Address  string `gorm:"size:255;comment:地址" json:"address"`
	Status   int8   `gorm:"size:1;comment:状态 0启用 1禁用" json:"status"`
	Age      int8   `gorm:"size:150;comment:年龄" json:"age"`
}

func (User) TableName() string {
	return "tb_user"
}
