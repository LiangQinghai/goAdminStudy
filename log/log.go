package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"goAdminStudy/config"
	"io"
	"os"
	"strings"
)

var logger *logrus.Logger

var DefaultLogFile = "./static/log.log"

func init() {

	logger = logrus.New()
	// 日志文件配置
	name := config.LogConfig.GetFileName()
	if name != "" {
		DefaultLogFile = name
	}
	var lastIndex int = 0
	for index, s := range DefaultLogFile {
		if s == '/' || s == '\\' {
			lastIndex = index
		}
	}
	path := DefaultLogFile[0:lastIndex]
	path = strings.ReplaceAll(path, "/", string(os.PathSeparator))
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, os.FileMode(06666))
		if err != nil {
			fmt.Printf("%v \n", err)
		}
	}
	file, err := os.OpenFile(DefaultLogFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.FileMode(0666))

	if err != nil {
		fmt.Println("log error: ", err)
		logger.SetOutput(os.Stdout)
	} else {
		//logger.Out = io.MultiWriter(file, os.Stdout)
		logger.SetOutput(io.MultiWriter(file, os.Stdout))
	}

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

func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args)
}

func Panicln(args ...interface{}) {
	logger.Panicln(args)
}
