package repository

import (
	"fmt"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"urlshortener/domain"
	"urlshortener/helper"
)

type User interface {
	Create(user *domain.User, ctx context.Context) string
	Auth(user domain.User) (domain.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (u *UserRepository) Create(user *domain.User, ctx context.Context) string {
	// hashing password before save to database
	user.Password = helper.HashAndSalt([]byte(user.Password))
	err := u.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	} else {
		return "nil"
	}
}

func (u *UserRepository) Auth(user domain.User) (domain.User, error) {
	// compare hashed password with password input
	var userDB domain.User
	err := u.db.Where("email = ?", user.Email).First(&userDB).Error
	if err != nil {
		return user, err
	}
	if !helper.ComparePassword(userDB.Password, []byte(user.Password)) {
		err = fmt.Errorf("password not match")
		// return 500 error
		return user, err
	}
	return userDB, err
}
