package repository

import (
	"gorm.io/gorm"
	"urlshortener/domain"
)

type User interface {
	Create(user domain.User) (domain.User, error)
	Auth(user domain.User) (domain.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (u *UserRepository) Create(user domain.User) (domain.User, error) {
	err := u.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserRepository) Auth(user domain.User) (domain.User, error) {
	err := u.db.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
