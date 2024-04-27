package conf

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"time"
)

func InitLogger() *zap.SugaredLogger {
	//定义日志输出的级别
	logModel := zapcore.DebugLevel
	if !viper.GetBool("logModel.develop") {
		logModel = zapcore.InfoLevel
	}
	core := zapcore.NewCore(getEncoder(), zapcore.NewMultiWriteSyncer(getWriterAsyncer(), zapcore.AddSync(os.Stdout)), logModel)
	return zap.New(core).Sugar()
}

// 对输出的格式进行定义
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	//配置输出时间字段的key  "time":"2023-03-05"
	encoderConfig.TimeKey = "Time"
	//将输出的等级的key格式化成大写
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	//对输出的时间进行格式化
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Local().Format("2006-01-02 15:04:05"))
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

// 输出的位置进行配置
func getWriterAsyncer() zapcore.WriteSyncer {
	//获取项目的分隔符
	stSeparator := string(filepath.Separator)
	//获取当前文件所在目录的绝对位置
	rootDir, _ := os.Getwd()
	//设置log文件所存放的位置
	logFilePath := rootDir + stSeparator + "log" + stSeparator + time.Now().Format("2006-01-02") + ".txt"

	//对日志文件进行分割
	lumberjackAsyncer := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    viper.GetInt("log.MaxSize"),
		MaxBackups: viper.GetInt("log.MaxBackups"),
		MaxAge:     viper.GetInt("log.MaxAge"),
		Compress:   false,
	}

	return zapcore.AddSync(lumberjackAsyncer)
}
