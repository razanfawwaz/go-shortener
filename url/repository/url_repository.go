package repository

import (
	"gorm.io/gorm"
	"urlshortener/domain"
)

type Url interface {
	GenerateUrl(url domain.Url) (domain.Url, error)
	FindUrl(short string) (string, error)
	GetAllUrl() ([]domain.Url, error)
	UpdateUrl(short string, url domain.Url) (domain.Url, error)
	DeleteUrl(url domain.Url) (domain.Url, error)
}

type UrlRepository struct {
	db *gorm.DB
}

func NewUrlRepository(db *gorm.DB) *UrlRepository {
	return &UrlRepository{db}
}

func (u *UrlRepository) GenerateUrl(url domain.Url) (domain.Url, error) {
	err := u.db.Create(&url).Error
	return url, err
}

func (u *UrlRepository) FindUrl(short string) (string, error) {
	var url domain.Url
	err := u.db.Where("short_url = ?", short).First(&url).Error
	// return long_url
	return url.LongUrl, err
}

func (u *UrlRepository) UpdateUrl(short string, url domain.Url) (domain.Url, error) {
	// update data by short id
	err := u.db.Where("short_url = ?", short).Updates(&url).Error
	return url, err
}

func (u *UrlRepository) GetAllUrl() ([]domain.Url, error) {
	var urls []domain.Url
	err := u.db.Find(&urls).Error
	return urls, err
}

func (u *UrlRepository) DeleteUrl(url domain.Url) (domain.Url, error) {
	err := u.db.Delete(&url).Error
	return url, err
}
