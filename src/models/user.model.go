package models

import (
	"errors"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/horlaarsco/bank/src/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Base struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type IUser struct {
	ID        uint   `gorm:"primary_key;auto_increment;" json:"id"`
	FirstName string `json:"first_name" valid:"required"`
	LastName  string `json:"last_name" valid:"required"`
	Email     string `json:"email" gorm:"unique;not null" valid:"required,email"`
	Password  string `json:"password" valid:"required,length(6|255)"`
	Base
	Token string `json:"token"`
}

type LoginDTO struct {
	Email    string `json:"email" valid:"required,email"`
	Password string `json:"password" valid:"required,length(6|255)"`
}

func (u *LoginDTO) ValidateBody() error {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		tt := strings.Split(strings.TrimRight(err.Error(), ";"), ";")
		return errors.New(tt[0])
	}
	return nil
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

func GetUserByEmail(email string, user *IUser) error {
	error := db.Where("email = ?", email).First(&user).Error
	if error != nil {
		if errors.Is(error, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return errors.New("an error occured")
	}
	return nil
}

func CreateUser(user *IUser) error {
	error := db.Create(&user).Error
	if error != nil {
		return error
	}
	return nil
}

func (u *IUser) ValidateBody() error {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		tt := strings.Split(strings.TrimRight(err.Error(), ";"), ";")
		return errors.New(tt[0])
	}
	return nil
}
