package usecase

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
	"urlshortener/domain"
)

type MockUserUsecase struct {
	mock.Mock
}

func (m *MockUserUsecase) Create(user *domain.User, ctx context.Context) string {
	args := m.Called(user, ctx)
	return args.String(0)
}

func (m *MockUserUsecase) Auth(user domain.User) (domain.User, error) {
	args := m.Called(user)
	return args.Get(0).(domain.User), args.Error(1)
}

type TestSuiteUserRepository struct {
	suite.Suite
	MockUserUsecase *MockUserUsecase
	userUsecase     UserUsecase
	ctx             context.Context
}

func (s *TestSuiteUserRepository) SetupTest() {
	s.MockUserUsecase = &MockUserUsecase{}
	s.userUsecase = NewUserUsecase(s.MockUserUsecase)
	s.ctx = context.Background()
}

func (s *TestSuiteUserRepository) TeardownTest() {
	s.MockUserUsecase = nil
	s.userUsecase = nil
	s.ctx = nil
}

func (s *TestSuiteUserRepository) TestCreateUser() {
	user := &domain.User{
		Email:    "razan@gmail.com",
		Password: "123456",
	}
	s.MockUserUsecase.On("Create", user, s.ctx).Return("1")
}

func (s *TestSuiteUserRepository) TestAuthUser() {
	user := domain.User{
		Email:    "razan@gmail.com",
		Password: "123456",
	}
	// check return value
	s.MockUserUsecase.On("Auth", user).Return(user, nil)
}

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(TestSuiteUserRepository))
}
