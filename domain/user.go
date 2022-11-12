package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string `json:"email" gorm:"unique"`
	Password     string `json:"password"`
	IsSubscribed bool   `default:"false"`
}

type UserResponse []UserResponse
