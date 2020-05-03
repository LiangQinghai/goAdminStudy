package models

import (
	"github.com/jinzhu/gorm"
	"goAdminStudy/datasource"
	"goAdminStudy/log"
)

func init() {
	if err := AutoMigration(datasource.DataSource); err != nil {
		log.Errorf("%v", err)
	}
}

func AutoMigration(db *gorm.DB) error {

	db.SingularTable(true)

	return db.AutoMigrate(
		new(User),
		new(KafkaConfig),
	).Error

}
