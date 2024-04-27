package common

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        string         `gorm:"type:varchar(200);not null;primary_key;comment:主键id" json:"id"`
	CreateBy  string         `gorm:"type:varchar(20);comment:创建者" json:"createBy"`
	CreatedAt time.Time      `gorm:"comment:创建时间" json:"createdAt"`
	UpdateBy  string         `gorm:"type:varchar(20);comment:修改者" json:"updateBy"`
	UpdatedAt time.Time      `gorm:"comment:更新时间" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:删除时间" json:"deletedAt"`
}

type BaseModel struct {
	ID        string         `gorm:"type:varchar(200);not null;primary_key;comment:主键id" json:"id"`
	CreatedAt time.Time      `gorm:"comment:创建时间" json:"createdAt"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:删除时间" json:"deletedAt"`
}
