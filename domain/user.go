package domain

import (
	"time"
)

type User struct {
	CreatedAt    time.Time
	ID           string `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	IsSubscribed bool   `json:"is_subscribed"`
	UpdatedAt    time.Time
	DeletedAt    time.Time `gorm:"index"`
}
