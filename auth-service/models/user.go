package models

import (
	"github.com/Abeldlp/reservation-service/auth-service/config"
	"gorm.io/gorm"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetAllUsers() []User {
	var users []User
	config.DB.Find(&users)
	return users
}

func CreateUser(user User) *gorm.DB {
	return config.DB.Create(&user)
}
