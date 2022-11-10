package usecase

import (
	"golang.org/x/net/context"
	"urlshortener/domain"
	"urlshortener/user/repository"
)

type UserUsecase interface {
	Create(user *domain.User, ctx context.Context) string
	Auth(user domain.User) (domain.User, error)
}

type userUsecase struct {
	userRepository repository.User
}

func NewUserUsecase(userRepo repository.User) *userUsecase {
	return &userUsecase{userRepo}
}

func (u *userUsecase) Create(user *domain.User, ctx context.Context) string {
	return u.userRepository.Create(user, ctx)
}

func (u *userUsecase) Auth(user domain.User) (domain.User, error) {
	return u.userRepository.Auth(user)
}
