package config

import "github.com/spf13/viper"

// 对外开放产量
var (
	DataSourceConf *dataSource

	JwtConfig *jwtConfig

	LogConfig *logConfig
)

func init() {
	viper.SetConfigName("application")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	initDataSource(viper.Sub("datasource"))

	initJwtConfig(viper.Sub("jwt"))

	initLogConfig(viper.Sub("log"))
}

func initDataSource(source *viper.Viper) {
	DataSourceConf = &dataSource{
		user:     source.GetString("user"),
		password: source.GetString("password"),
		host:     source.GetString("host"),
		port:     source.GetInt("port"),
		dataBase: source.GetString("database"),
	}
}

func initJwtConfig(source *viper.Viper) {
	JwtConfig = &jwtConfig{
		secret: source.GetString("secret"),
	}
}

func initLogConfig(source *viper.Viper) {
	LogConfig = &logConfig{
		fileName: source.GetString("fileName"),
	}
}
