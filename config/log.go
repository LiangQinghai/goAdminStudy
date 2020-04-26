package config

// 日志文件配置
type logConfig struct {
	fileName string
}

func (l *logConfig) GetFileName() string {
	return l.fileName
}
