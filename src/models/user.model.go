package models

import (
	"errors"

	"github.com/horlaarsco/bank/src/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type IUser struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique;not null"`
	Password  string `json:"password"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&IUser{})
}

func GetUser(id uint, user *IUser) string {
	error := db.First(&user, id).Error
	if error != nil {
		if errors.Is(error, gorm.ErrRecordNotFound) {
			return "User not found"
		}
		return "An Error Occured"
	}
	return ""
}

func CreateUser(user IUser) string {
	error := db.Create(&user).Error
	if error != nil {
		return error.Error()
	}
	return ""
}
