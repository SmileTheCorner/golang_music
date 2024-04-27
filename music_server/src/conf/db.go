package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"go_study/src/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

func InitDb() {
	logModel := logger.Info
	if !viper.GetBool("logModel.develop") {
		logModel = logger.Error
	}
	db, err := gorm.Open(mysql.Open(viper.GetString("db.mysql")), &gorm.Config{
		//配置生成表名的策略
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tb_",
			SingularTable: true,
		},
		//设置控制台输出sql语句的日志级别
		Logger: logger.Default.LogMode(logModel),
	})
	//db.Set("gorm:table_options", "COMMENT='角色菜单表'").AutoMigrate(entity.RoleMenu{})
	//db.Set("gorm:table_options", "COMMENT='部门表'").AutoMigrate(entity.Department{})
	//db.Set("gorm:table_options", "COMMENT='字典类型表'").AutoMigrate(entity.DictType{})
	//db.Set("gorm:table_options", "COMMENT='字典数值表'").AutoMigrate(entity.DictData{})
	//db.Set("gorm:table_options", "COMMENT='菜单表'").AutoMigrate(entity.Menu{})
	//db.Set("gorm:table_options", "COMMENT='菜单元信息表'").AutoMigrate(sys2.MenuMeta{})
	//db.Set("gorm:table_options", "COMMENT='用户表'").AutoMigrate(model.User{})
	//db.Set("gorm:table_options", "COMMENT='歌手表'").AutoMigrate(model.Singer{})
	//db.Set("gorm:table_options", "comment='菜单表'").AutoMigrate(model.Menu{})
	//db.Set("gorm:table_options", "comment='资源表'").AutoMigrate(model.Resource{})
	//db.Set("gorm:table_options", "comment='角色表'").AutoMigrate(model.Role{})
	//db.Set("gorm:table_options", "comment='角色菜单表'").AutoMigrate(model.RoleMenu{})
	//db.Set("gorm:table_options", "comment='角色资源表'").AutoMigrate(model.RoleResource{})
	//db.Set("gorm:table_options", "comment='角色用户表'").AutoMigrate(model.RoleUser{})
	//db.Set("gorm:table_options", "comment='歌词表'").AutoMigrate(entity.Lyric{})
	//db.Set("gorm:table_options", "comment='代码生成业务表'").AutoMigrate(entity.GenTable{})
	//db.Set("gorm:table_options", "comment='代码生成业务字段表'").AutoMigrate(entity.GenTableColumn{})

	sqlDb, _ := db.DB()

	sqlDb.SetMaxIdleConns(viper.GetInt("db.maxIdleConns"))
	sqlDb.SetMaxOpenConns(viper.GetInt("db.maxOpenConns"))
	sqlDb.SetConnMaxLifetime(time.Hour)

	if err != nil {
		panic(fmt.Sprintf("Connect Database Error : %s", err.Error()))
	}

	global.DB = db
}
