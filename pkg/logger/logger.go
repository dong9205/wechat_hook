package logger

import (
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logger zerolog.Logger
	l      *lumberjack.Logger
)

func InitLogger() {
	l := &lumberjack.Logger{
		Filename:   "./logs/wechat_hook.log",
		MaxSize:    500, // megabytes
		MaxBackups: 30,
		MaxAge:     15,   // days
		Compress:   true, // disabled by default
	}
	multi := zerolog.MultiLevelWriter(zerolog.NewConsoleWriter(), l)
	logger = zerolog.New(multi).With().Timestamp().Logger()
}

func GetLogger() zerolog.Logger {
	return logger
}
