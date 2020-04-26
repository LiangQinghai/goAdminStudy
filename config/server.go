package config

// 服务器配置
type serverConfig struct {
	port int
	host string
}

func (s *serverConfig) GetPort() int {
	return s.port
}

func (s *serverConfig) GetHost() string {
	return s.host
}
