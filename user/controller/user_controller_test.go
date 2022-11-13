package controller

// testing for controller

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"golang.org/x/net/context"
	"net/http"
	"net/http/httptest"
	"testing"
	"urlshortener/domain"
)

type MockUserUsecase struct {
	mock.Mock
}

func (m *MockUserUsecase) Create(user *domain.User, ctx context.Context) error {
	args := m.Called(user, ctx)
	return args.Error(0)
}

func (m *MockUserUsecase) Auth(user domain.User) (domain.User, error) {
	args := m.Called(user)
	return args.Get(0).(domain.User), args.Error(1)
}

type TestSuiteUserController struct {
	suite.Suite
}

func (s *TestSuiteUserController) TestCreate() {
	e := echo.New()
	user := domain.User{
		Email:    "mail@gmail.com",
		Password: "password",
	}
	{
		req := httptest.NewRequest(http.MethodPost, "/register", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("")
		c.Set("user", user)

		mockUserUsecase := new(MockUserUsecase)
		mockUserUsecase.On("Create", &user, mock.Anything).Return(nil)
	}
}

// test auth
func (s *TestSuiteUserController) TestAuth() {
	e := echo.New()
	user := domain.User{
		Email:    "mail@gmail.com",
		Password: "password",
	}
	{
		req := httptest.NewRequest(http.MethodPost, "/login", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("")
		c.Set("user", user)
		mockUserUsecase := new(MockUserUsecase)
		mockUserUsecase.On("Auth", user).Return("", nil)
	}
}

func TestUserController(t *testing.T) {
	suite.Run(t, new(TestSuiteUserController))
	t.Log()
}
