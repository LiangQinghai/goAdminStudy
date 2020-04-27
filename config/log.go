package config

import (
	"github.com/sirupsen/logrus"
	"strings"
)

// 日志文件配置
type logConfig struct {
	fileName string
	level    string
}

func (l *logConfig) GetFileName() string {
	return l.fileName
}

// 日志级别
func (l *logConfig) GetLevel() logrus.Level {
	switch strings.ToLower(l.level) {
	case "trace":
		return logrus.TraceLevel
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel
	case "panic":
		return logrus.PanicLevel
	default:
		return logrus.InfoLevel
	}
}
