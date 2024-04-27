package conf

import (
	"github.com/spf13/viper"
)

func InitConfig() {
	//对viper做基本的配置
	//设置配置文件的名称
	viper.SetConfigName("setting")
	//设置配置文件的类型
	viper.SetConfigType("yml")
	//设置配置文件的路径
	viper.AddConfigPath("./src/conf/")
	//读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		panic("Load Config Error:" + err.Error())
	}
}
