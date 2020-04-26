package config

// jwt配置
type jwtConfig struct {
	secret string
}

func (j *jwtConfig) GetSecret() string {
	return j.secret
}
