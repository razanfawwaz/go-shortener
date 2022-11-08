package usecase

import (
	"urlshortener/domain"
	"urlshortener/url/repository"
)

type UseCase interface {
	GenerateUrl(url domain.Url) (domain.Url, error)
	FindUrl(short string) (string, error)
	UpdateUrl(short string, id int, url domain.Url) (domain.Url, error)
	DeleteUrl(url domain.Url) (domain.Url, error)
	GetAllUrl() ([]domain.Url, error)
	UserUrl(id int) ([]domain.Url, error)
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

func (u *urlUsecase) DeleteUrl(url domain.Url) (domain.Url, error) {
	return u.urlRepository.DeleteUrl(url)
}

func (u *urlUsecase) UserUrl(id int) ([]domain.Url, error) {
	return u.urlRepository.UserUrl(id)
}
