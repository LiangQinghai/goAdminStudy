package models

import (
	"github.com/jinzhu/gorm"
	"goAdminStudy/datasource"
	"goAdminStudy/utils"
	"time"
)

type KafkaConfig struct {
	Name string `json:"name" gorm:"type:varchar(100);not null;unique;comment:'kafka名称'" binding:"required"`

	Broker string `json:"broker" gorm:"type:varchar(200);not null;comment:'kafka broker'" binding:"required"`

	Desc string `json:"desc" gorm:"type:text;comment:'描述'"`

	Status int `json:"status" gorm:"type:int;not null;default:0;comment:'状态'"`

	BaseModule
}

func (KafkaConfig) TableName() string {
	return "T_KAFKA_CONFIG"
}

func (k *KafkaConfig) db() *gorm.DB {
	return datasource.DataSource.Table(k.TableName())
}

func (k *KafkaConfig) Save() error {

	k.UpdateTime = time.Now()
	k.CreatTime = time.Now()
	k.Status = StatusEnable

	save := k.db().Save(&k)

	return save.Error

}

func (k *KafkaConfig) GetByPage(page *utils.Page) (kafkaConfigs []KafkaConfig, err error) {

	db := k.db().Select("*")

	if page.KeyWorld != "" {
		db = db.Where("name LIKE %?%", page.KeyWorld).Or("broker LIKE %?%", page.KeyWorld).Or("desc LIKE %?%", page.KeyWorld)
	}

	err = db.Offset((page.GetPageNo() - 1) * page.GetPageCount()).Limit(page.GetPageCount()).Find(&kafkaConfigs).Error

	db.Count(&(page.TotalCount))

	return

}
