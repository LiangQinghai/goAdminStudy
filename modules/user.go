package modules

import (
	"github.com/jinzhu/gorm"
	"goAdminStudy/datasource"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// 用户
type User struct {
	UserName string `json:"user_name";gorm:"UNIQUE;Type:varchar(100);not_null"`
	Password string `json:"password";gorm:"type:varchar(100);not_null"`
	Status   int    `json:"status";gorm:"type:tinyint(2);not_null;DEFAULT:0;comment:'0->enable, 1->disable, 2->delete'"`
	BaseModule
}

func (User) TableName() string {
	return "T_USER"
}

func (u *User) db() *gorm.DB {
	return datasource.DataSource.Table(u.TableName())
}

// 创建
func (u *User) Create() error {

	u.CreatTime = time.Now()
	u.UpdateTime = time.Now()
	u.Status = StatusEnable

	// 加密
	if password, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost); err != nil {
		return err
	} else {
		u.Password = string(password)
	}

	result := u.db().Create(&u)

	if err := result.Error; err != nil {
		return err
	}

	return nil

}

func (u *User) Update() error {

	u.UpdateTime = time.Now()
	if err := u.db().Update(&u).Error; err != nil {
		return err
	}
	return nil

}

func (u *User) GetById() (user User, err error) {

	result := u.db().Where("id = ?", u.Id)

	if err = result.Error; err != nil {
		return
	}

	err = result.First(&user).Error

	return

}

func (u *User) GetByPage(pageSize int, pageNo int) (users []User, pageCount int, err error) {

	db := u.db()

	if u.UserName != "" {
		db.Where("user_name like", "%"+u.UserName+"%")
	}

	err = db.Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&users).Count(&pageCount).Error

	return

}
