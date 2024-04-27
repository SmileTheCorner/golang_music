package conf

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go_study/src/global"
	"time"
)

// InitRedis 初始化redis
func InitRedis() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})
	_, err := redisClient.Ping(context.Background()).Result()

	if err != nil {
		global.Logger.Error(fmt.Sprintf("Connect Redis Fail : %s", err.Error()))
		panic(fmt.Sprintf("Connect Redis Fail : %s", err.Error()))
	}
	global.RedisClient = redisClient
}

// RsSetKey 设置k - v
func RsSetKey(key string, value interface{}, expire ...int8) {
	if expire[0] == 0 {
		expire[0] = 60
	}
	err := global.RedisClient.Set(context.Background(), key, value, time.Duration(expire[0])*time.Minute).Err()
	if err != nil {
		global.Logger.Error(fmt.Sprintf("Redis Set Key Fail : %s", err.Error()))
		return
	}
	global.Logger.Error(fmt.Sprintf("Redis Set Key Success : %s", key))
}

// RsGetKey 通过key获取对应的值
func RsGetKey(key string) any {
	val, err := global.RedisClient.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return fmt.Sprintf("%s Does Not Exist", key)
	} else if err != nil {
		return fmt.Sprintf("Get %s Fail : %s", key, err.Error())
	} else {
		return val
	}
}
