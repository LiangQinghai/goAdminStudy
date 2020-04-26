package datasource

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"goAdminStudy/config"
)

var DataSource *gorm.DB

func init() {

	conf := config.DataSourceConf
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.GetUser(), conf.GetPassword(), conf.GetHost(), conf.GetPort(), conf.GetDataBase())
	var err error = nil
	if DataSource, err = gorm.Open("mysql", url); err != nil {
		panic(err)
	}

	DataSource.LogMode(true).Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")

}
