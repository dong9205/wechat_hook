package logger

import (
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logger zerolog.Logger
	l      *lumberjack.Logger
)

func InitLogger() {
	var (
		logFilename = viper.GetString("logger.filename")
		maxSize     = viper.GetInt("logger.max_size")
		maxAge      = viper.GetInt("logger.max_age")
		maxBack     = viper.GetInt("logger.max_backups")
		compress    = viper.GetBool("logger.compress")
		level       = viper.GetString("logger.level")
	)
	if logFilename == "" {
		logFilename = "./logs/wechat_hook.log"
	}
	if maxSize == 0 {
		maxSize = 500
	}
	if maxAge == 0 {
		maxAge = 15
	}
	if maxBack == 0 {
		maxBack = 30
	}
	l := &lumberjack.Logger{
		Filename:   logFilename,
		MaxSize:    maxSize, // megabytes
		MaxAge:     maxAge,  // days
		MaxBackups: maxBack,
		Compress:   compress, // disabled by default
	}
	zlevel, err := zerolog.ParseLevel(level)
	if err != nil {
		zlevel = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(zlevel)
	multi := zerolog.MultiLevelWriter(zerolog.NewConsoleWriter(), l)
	logger = zerolog.New(multi).With().Timestamp().Logger()
}

func GetLogger() zerolog.Logger {
	return logger
}
