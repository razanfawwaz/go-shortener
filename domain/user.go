package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string `json:"email"`
	Password     string `json:"password"`
	IsSubscribed bool   `json:"is_subscribed"`
}
