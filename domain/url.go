package domain

import (
	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	LongUrl   string `json:"long_url"`
	ShortUrl  string `json:"short_url" gorm:"unique"`
	Clicks    int
	TotalEdit int
	UserID    uint
	User      User
}
