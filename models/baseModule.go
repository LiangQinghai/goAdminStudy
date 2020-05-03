package models

import "time"

//
type BaseModule struct {
	Id         int32     `json:"id" gorm:"primary_key;AUTO_INCREMENT;type:bigint(20)"`
	CreatTime  time.Time `json:"creat_time" gorm:"Type:DATETIME"`
	UpdateTime time.Time `json:"update_time" gorm:"Type:DATETIME"`
}
