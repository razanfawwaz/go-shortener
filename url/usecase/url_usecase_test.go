package usecase

import (
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
	"urlshortener/domain"
)

type MockUrlUsecase struct {
	mock.Mock
}

func (m *MockUrlUsecase) GenerateUrl(url domain.Url) (domain.Url, error) {
	args := m.Called(url)
	return args.Get(0).(domain.Url), args.Error(1)
}

func (m *MockUrlUsecase) FindUrl(short string) (domain.Url, error) {
	args := m.Called(short)
	return args.Get(0).(domain.Url), args.Error(1)
}

func (m *MockUrlUsecase) GetAllUrl() ([]domain.Url, error) {
	args := m.Called()
	return args.Get(0).([]domain.Url), args.Error(1)
}

func (m *MockUrlUsecase) DeleteUrl(short string) error {
	args := m.Called(short)
	return args.Error(0)
}

func (m *MockUrlUsecase) UpdateUrl(short string, url domain.Url) (domain.Url, error) {
	args := m.Called(short, url)
	return args.Get(0).(domain.Url), args.Error(1)
}

func (m *MockUrlUsecase) UserUrl(id int) ([]domain.Url, error) {
	args := m.Called(id)
	return args.Get(0).([]domain.Url), args.Error(1)
}

func (m *MockUrlUsecase) ExpiredUrl(short string) (bool, error) {
	args := m.Called(short)
	return args.Get(0).(bool), args.Error(1)
}

func (m *MockUrlUsecase) SubsStatus(id int) (bool, error) {
	args := m.Called(id)
	return args.Get(0).(bool), args.Error(1)
}

func (m *MockUrlUsecase) GetUrlByUser(id int) ([]domain.Url, error) {
	args := m.Called(id)
	return args.Get(0).([]domain.Url), args.Error(1)
}

type TestSuiteUrlRepository struct {
	suite.Suite
	urlUsecase     *MockUrlUsecase
	mockUrlUsecase *MockUrlUsecase
}

func (s *TestSuiteUrlRepository) SetupTest() {
	s.mockUrlUsecase = new(MockUrlUsecase)
	s.urlUsecase = s.mockUrlUsecase
}

func (s *TestSuiteUrlRepository) TestGenerateUrl() {
	s.mockUrlUsecase.On("GenerateUrl", mock.AnythingOfType("domain.Url")).Return(domain.Url{}, nil).Once()
	_, err := s.urlUsecase.GenerateUrl(domain.Url{})
	s.NoError(err)
}

func (s *TestSuiteUrlRepository) TestFindUrl() {
	s.mockUrlUsecase.On("FindUrl", mock.AnythingOfType("string")).Return(domain.Url{}, nil).Once()
	_, err := s.urlUsecase.FindUrl("")
	s.NoError(err)
}

func (s *TestSuiteUrlRepository) TestGetAllUrl() {
	s.mockUrlUsecase.On("GetAllUrl").Return([]domain.Url{}, nil).Once()
	_, err := s.urlUsecase.GetAllUrl()
	s.NoError(err)
}

func (s *TestSuiteUrlRepository) TestDeleteUrl() {
	s.mockUrlUsecase.On("DeleteUrl", mock.AnythingOfType("string")).Return(nil).Once()
	err := s.urlUsecase.DeleteUrl("")
	s.NoError(err)
}

func (s *TestSuiteUrlRepository) TestUpdateUrl() {
	s.mockUrlUsecase.On("UpdateUrl", mock.AnythingOfType("string"), mock.AnythingOfType("domain.Url")).Return(domain.Url{}, nil).Once()
	_, err := s.urlUsecase.UpdateUrl("", domain.Url{})
	s.NoError(err)
}

func (s *TestSuiteUrlRepository) TestUserUrl() {
	s.mockUrlUsecase.On("UserUrl", mock.AnythingOfType("int")).Return([]domain.Url{}, nil).Once()
	_, err := s.urlUsecase.UserUrl(0)
	s.NoError(err)
}

func (s *TestSuiteUrlRepository) TestExpiredUrl() {
	s.mockUrlUsecase.On("ExpiredUrl", mock.AnythingOfType("string")).Return(false, nil).Once()
	_, err := s.urlUsecase.ExpiredUrl("")
	s.NoError(err)
}

func (s *TestSuiteUrlRepository) TestSubsStatus() {
	s.mockUrlUsecase.On("SubsStatus", mock.AnythingOfType("int")).Return(false, nil).Once()
	_, err := s.urlUsecase.SubsStatus(0)
	s.NoError(err)
}

func (s *TestSuiteUrlRepository) TestGetUrlByUser() {
	s.mockUrlUsecase.On("GetUrlByUser", mock.AnythingOfType("int")).Return([]domain.Url{}, nil).Once()
	_, err := s.urlUsecase.GetUrlByUser(0)
	s.NoError(err)
}

func TestUrlUsecase(t *testing.T) {
	suite.Run(t, new(TestSuiteUrlRepository))
}
