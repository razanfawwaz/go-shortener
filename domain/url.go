package domain

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:char(36);primary_key;default:uuid_generate_v4()"`
	LongUrl   string    `json:"long_url"`
	ShortUrl  string    `json:"short_url" gorm:"unique"`
	Clicks    int
	TotalEdit int
}
