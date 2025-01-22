package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model

	FirstName string     `gorm:"first_name; type:varchar(100)"`
	LastName  string     `gorm:"last_name; type:varchar(100)"`
	Email     string     `gorm:"email; type:varchar(100)"`
	Username  string     `gorm:"username; type:varchar(100)"`
	Password  string     `gorm:"password; type:varchar(100)"`
	Gender    string     `gorm:"gender; type:varchar(100)"`
	Age       uint8      `gorm:"age; type:tinyint(4) unsigned"`
	Birthday  *time.Time `gorm:"birthday"`
	Address   string     `gorm:"address; type:varchar(100)"`
	City      string     `gorm:"city; type:varchar(100)"`
	State     string     `gorm:"state; type:varchar(100)"`
	Country   string     `gorm:"country; type:varchar(100)"`
}

func (Users) TableName() string {
	return "users"
}
