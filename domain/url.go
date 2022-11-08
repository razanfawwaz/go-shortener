package domain

import (
	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	LongUrl   string `json:"long_url"`
	ShortUrl  string `json:"short_url" gorm:"unique"`
	Clicks    int    `default:"0"`
	TotalEdit int
	UserID    int
}
