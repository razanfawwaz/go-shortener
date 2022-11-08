package repository

import (
	"fmt"
	"gorm.io/gorm"
	"urlshortener/domain"
)

type Url interface {
	GenerateUrl(url domain.Url) (domain.Url, error)
	FindUrl(short string) (string, error)
	GetAllUrl() ([]domain.Url, error)
	UpdateUrl(short string, id int, url domain.Url) (domain.Url, error)
	DeleteUrl(short string, id int, url domain.Url) error
	UserUrl(id int) ([]domain.Url, error)
	ExpiredUrl(short string) (bool, error)
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

func (u *UrlRepository) ExpiredUrl(short string) (bool, error) {
	var urls []domain.Url
	// check expired date using short_url and expired_at
	date := u.db.Where("short_url = ? AND expired_at > NOW()", short).Find(&urls)
	// If date expired > date now, the date not already expired so return false
	if date.RowsAffected != 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func (u *UrlRepository) FindUrl(short string) (string, error) {
	var url domain.Url

	err := u.db.Where("short_url = ?", short).First(&url).Error
	expired, _ := u.ExpiredUrl(short)

	// If err == nil & expired == false, count clicks & return url. If not, return error
	if err == nil && !expired {
		u.db.Model(&url).Update("clicks", gorm.Expr("clicks + ?", 1))
	} else if err == nil && expired {
		err = fmt.Errorf("URL Time Expired")
		return url.LongUrl, err
	} else {
		err = gorm.ErrRecordNotFound
		return url.LongUrl, err
	}
	return url.LongUrl, err
}

func (u *UrlRepository) UpdateUrl(short string, id int, url domain.Url) (domain.Url, error) {
	err := u.db.Model(&url).Where("short_url = ? AND user_id = ?", short, id).Updates(url).Error
	return url, err
}

func (u *UrlRepository) GetAllUrl() ([]domain.Url, error) {
	var urls []domain.Url
	err := u.db.Find(&urls).Error
	return urls, err
}

func (u *UrlRepository) DeleteUrl(short string, id int, url domain.Url) error {
	err := u.db.Where("short_url = ? AND user_id = ?", short, id).Delete(&url).Error
	return err
}

func (u *UrlRepository) UserUrl(id int) ([]domain.Url, error) {
	var urls []domain.Url
	err := u.db.Where("user_id = ?", id).Find(&urls).Error
	return urls, err
}
