package repository

import (
	"fmt"
	"gorm.io/gorm"
	"urlshortener/domain"
	"urlshortener/helper"
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
	// hashing password before save to database
	user.Password = helper.HashAndSalt([]byte(user.Password))
	err := u.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
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
	fmt.Println("password match", userDB.Password, user.Password)
	return userDB, err
}
