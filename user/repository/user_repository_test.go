package repository

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"regexp"
	"testing"
	"urlshortener/config"
)

type TestSuiteUserRepository struct {
	suite.Suite
	Mock           sqlmock.Sqlmock
	userRepository User
	ctx            context.Context
}

func (s *TestSuiteUserRepository) SetupTest() {
	dbMock, mock, err := sqlmock.New()
	s.NoError(err)

	DB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	s.NoError(err)

	DB = config.DB

	s.Mock = mock
	s.userRepository = NewUserRepository(DB)
	s.ctx = context.Background()
}

func (s *TestSuiteUserRepository) TeardownTest() {
	s.Mock = nil
	s.userRepository = nil
	s.ctx = nil
}

func (s *TestSuiteUserRepository) TestCreateUser() {
	for _, tt := range []struct {
		Name        string
		Query       string
		Err         error
		ExpectedErr error
	}{
		{
			Name:  "Success",
			Query: "INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`email`,`password`,`is_subscribed`) VALUES (?,?,?,?,?,?)",
		},
		{
			Name:        "Error from DB",
			Query:       "INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`email`,`password`,`is_subscribed`) VALUES (?,?,?,?,?,?)",
			Err:         errors.New("error"),
			ExpectedErr: errors.New("error"),
		},
	} {
		s.SetupTest()
		s.Run(tt.Name, func() {
			s.Mock.ExpectBegin()
			if tt.Err != nil {
				s.Mock.ExpectExec(regexp.QuoteMeta(tt.Query)).WillReturnError(tt.Err)
				s.Mock.ExpectRollback()
			} else {
				s.Mock.ExpectExec(regexp.QuoteMeta(tt.Query)).WillReturnResult(sqlmock.NewResult(1, 1))
				s.Mock.ExpectCommit()
			}

		})
		s.TeardownTest()
	}
}

func (s *TestSuiteUserRepository) TestAuthUser() {
	for _, tt := range []struct {
		Name        string
		Query       string
		Err         error
		ExpectedErr error
	}{
		{
			Name:  "Success",
			Query: "SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL AND ((email = ?)) ORDER BY `users`.`id` LIMIT 1",
		},
		{
			Name:        "Generic Error from DB",
			Query:       "SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL AND ((email = ?)) ORDER BY `users`.`id` LIMIT 1",
			Err:         errors.New("generic error"),
			ExpectedErr: errors.New("generic error"),
		},
	} {
		s.SetupTest()
		s.Run(tt.Name, func() {
			s.Mock.ExpectBegin()
			if tt.Err != nil {
				s.Mock.ExpectExec(regexp.QuoteMeta(tt.Query)).WillReturnError(tt.Err)
				s.Mock.ExpectRollback()
			} else {
				s.Mock.ExpectExec(regexp.QuoteMeta(tt.Query)).WillReturnResult(sqlmock.NewResult(1, 1))
				s.Mock.ExpectCommit()
			}

		})
		s.TeardownTest()
	}
}

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(TestSuiteUserRepository))
}
