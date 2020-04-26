package config

// 数据源配置
type dataSource struct {
	user     string
	password string
	host     string
	port     int
	dataBase string
}

func (d *dataSource) GetUser() string {
	return d.user
}

func (d *dataSource) GetPassword() string {
	return d.password
}

func (d *dataSource) GetHost() string {
	return d.host
}

func (d *dataSource) GetPort() int {
	return d.port
}

func (d *dataSource) GetDataBase() string {
	return d.dataBase
}
