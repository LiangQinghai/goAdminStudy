package datasource

import (
	"fmt"
	"github.com/LiangQinghai/goAdminStudy/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DataSource *gorm.DB

func init() {
	//"root:liangqinghai@tcp(127.0.0.1:3306)/test"
	conf := config.DataSourceConf
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.User, conf.Password, conf.Host, conf.Port, conf.DataBase)
	var err error = nil
	if DataSource, err = gorm.Open("mysql", url); err != nil {
		panic(err)
	}

	DataSource.LogMode(true).Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")

}
