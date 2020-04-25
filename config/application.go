package config

import "github.com/spf13/viper"

type DataSource struct {
	User     string
	Password string
	Host     string
	Port     int
	DataBase string
}

var DataSourceConf *DataSource

func init() {
	viper.SetConfigName("application")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	assembleDataSource(viper.Sub("datasource"))
}

func assembleDataSource(source *viper.Viper) {
	DataSourceConf = &DataSource{
		User:     source.GetString("user"),
		Password: source.GetString("password"),
		Host:     source.GetString("host"),
		Port:     source.GetInt("port"),
		DataBase: source.GetString("database"),
	}
}
