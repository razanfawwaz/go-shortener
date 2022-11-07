package usecase

import (
	"urlshortener/domain"
	"urlshortener/user/repository"
)

type UserUsecase interface {
	Create(user domain.User) (domain.User, error)
	Auth(user domain.User) (domain.User, error)
}

type userUsecase struct {
	userRepository repository.User
}

func NewUserUsecase(userRepo repository.User) *userUsecase {
	return &userUsecase{userRepo}
}

func (u *userUsecase) Create(user domain.User) (domain.User, error) {
	return u.userRepository.Create(user)
}

func (u *userUsecase) Auth(user domain.User) (domain.User, error) {
	return u.userRepository.Auth(user)
}
