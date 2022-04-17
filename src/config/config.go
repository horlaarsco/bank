package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	var err error
	dsn := "root:fnd143430078@/bank?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
