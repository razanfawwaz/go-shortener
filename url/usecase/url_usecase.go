package usecase

import (
	"urlshortener/domain"
	"urlshortener/url/repository"
)

type UseCase interface {
	GenerateUrl(url domain.Url) (domain.Url, error)
	UpdateUrl(short string, id int, url domain.Url) (domain.Url, error)
	DeleteUrl(short string, id int, url domain.Url) error
	FindUrl(short string) (string, error)
	GetAllUrl() ([]domain.Url, error)
	UserUrl(id int) ([]domain.Url, error)
	ExpiredUrl(short string) (bool, error)
}

type urlUsecase struct {
	urlRepository repository.Url
}

func NewUrlUsecase(urlRepo repository.Url) *urlUsecase {
	return &urlUsecase{urlRepo}
}

func (u *urlUsecase) GenerateUrl(url domain.Url) (domain.Url, error) {
	return u.urlRepository.GenerateUrl(url)
}

func (u *urlUsecase) FindUrl(short string) (string, error) {
	return u.urlRepository.FindUrl(short)
}

func (u *urlUsecase) GetAllUrl() ([]domain.Url, error) {
	return u.urlRepository.GetAllUrl()
}

func (u *urlUsecase) UpdateUrl(short string, id int, url domain.Url) (domain.Url, error) {
	return u.urlRepository.UpdateUrl(short, id, url)
}

func (u *urlUsecase) DeleteUrl(short string, id int, url domain.Url) error {
	return u.urlRepository.DeleteUrl(short, id, url)
}

func (u *urlUsecase) UserUrl(id int) ([]domain.Url, error) {
	return u.urlRepository.UserUrl(id)
}

func (u *urlUsecase) ExpiredUrl(short string) (bool, error) {
	return u.urlRepository.ExpiredUrl(short)
}
