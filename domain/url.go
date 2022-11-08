package domain

import (
	"gorm.io/gorm"
	"time"
)

type Url struct {
	gorm.Model
	LongUrl   string    `json:"long_url"`
	ShortUrl  string    `json:"short_url" gorm:"unique"`
	Clicks    int       `default:"0"`
	ExpiredAt time.Time `json:"expired_at" gorm:"default:NULL"`
	TotalEdit int
	UserID    int
}
