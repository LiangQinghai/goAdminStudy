package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"goAdminStudy/config"
	"io"
	"os"
)

var logger *logrus.Logger

func init() {

	logger = logrus.New()
	// 日志文件配置
	file, err := os.OpenFile(config.LogConfig.GetFileName(), os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		fmt.Println("log error: ", err)
	}

	//logger.Out = io.MultiWriter(file, os.Stdout)
	logger.SetOutput(io.MultiWriter(file, os.Stdout))
	//
	logger.SetLevel(config.LogConfig.GetLevel())
	//
	logger.SetFormatter(&logrus.TextFormatter{})

}

func Debug(args ...interface{}) {
	logger.Debug(args)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args)
}

func Debugnln(args ...interface{}) {
	logger.Debugln(args)
}

func Info(args ...interface{}) {
	logger.Info(args)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args)
}

func Infoln(args ...interface{}) {
	logger.Infoln(args)
}

func Warn(args ...interface{}) {
	logger.Warn(args)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args)
}

func Warnln(args ...interface{}) {
	logger.Warnln(args)
}

func Error(args ...interface{}) {
	logger.Error(args)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args)
}

func Errorln(args ...interface{}) {
	logger.Errorln(args)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args)
}

func Fatalf(format string, args ...interface{}) {
	logger.Fatal(format, args)
}

func Fatalln(args ...interface{}) {
	logger.Fatalln(args)
}

func Panic(args ...interface{}) {
	logger.Panic(args)
}

func Panicf(args ...interface{}) {
	logger.Panic(args)
}

func Panic(args ...interface{}) {
	logger.Panic(args)
}
