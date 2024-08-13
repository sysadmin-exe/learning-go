package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// TODO: Make DB credenials configurable via environment variables
func Connect() {
	d, err := gorm.Open("mysql", "<db_user_name>:<db_user_password>@tcp(<db_hostname>:<db_port>)/<db_name>?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
