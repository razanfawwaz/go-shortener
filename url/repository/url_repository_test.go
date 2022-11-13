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

type TestSuiteUrlRepository struct {
	suite.Suite
	Mock          sqlmock.Sqlmock
	urlRepository Url
	ctx           context.Context
}

func (s *TestSuiteUrlRepository) SetupTest() {
	dbMock, mock, err := sqlmock.New()
	s.NoError(err)

	DB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	s.NoError(err)

	DB = config.DB

	s.Mock = mock
	s.urlRepository = NewUrlRepository(DB)
	s.ctx = context.Background()
}

func (s *TestSuiteUrlRepository) TeardownTest() {
	s.Mock = nil
	s.urlRepository = nil
	s.ctx = nil
}

func (s *TestSuiteUrlRepository) TestCreateUrl() {
	for _, tt := range []struct {
		Name        string
		Query       string
		Err         error
		ExpectedErr error
	}{
		{
			Name:  "Success",
			Query: "INSERT INTO `urls` (`created_at`,`updated_at`,`deleted_at`,`url`,`short_url`,`user_id`) VALUES (?,?,?,?,?,?)",
		},
		{
			Name:        "Error from DB",
			Query:       "INSERT INTO `urls` (`created_at`,`updated_at`,`deleted_at`,`url`,`short_url`,`user_id`) VALUES (?,?,?,?,?,?)",
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
	}
}

func (s *TestSuiteUrlRepository) TestGetUrl() {
	for _, tt := range []struct {
		Name        string
		Query       string
		Err         error
		ExpectedErr error
	}{
		{
			Name:  "Success",
			Query: "SELECT * FROM `urls` WHERE `urls`.`deleted_at` IS NULL AND ((`urls`.`short_url` = ?)) ORDER BY `urls`.`id` LIMIT 1",
		},
		{
			Name:        "Error from DB",
			Query:       "SELECT * FROM `urls` WHERE `urls`.`deleted_at` IS NULL AND ((`urls`.`short_url` = ?)) ORDER BY `urls`.`id` LIMIT 1",
			Err:         errors.New("error"),
			ExpectedErr: errors.New("error"),
		},
	} {
		s.SetupTest()
		s.Run(tt.Name, func() {
			s.Mock.ExpectBegin()
			if tt.Err != nil {
				s.Mock.ExpectQuery(regexp.QuoteMeta(tt.Query)).WillReturnError(tt.Err)
				s.Mock.ExpectRollback()
			} else {
				s.Mock.ExpectQuery(regexp.QuoteMeta(tt.Query)).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "url", "short_url", "user_id"}).AddRow(1, "2021-05-03 15:34:07", "2021-05-03 15:34:07", nil, "https://www.google.com/", "google", 1))
				s.Mock.ExpectCommit()
			}
		})
	}
}

func (s *TestSuiteUrlRepository) TestGetUrlByUserId() {
	for _, tt := range []struct {
		Name        string
		Query       string
		Err         error
		ExpectedErr error
	}{
		{
			Name:  "Success",
			Query: "SELECT * FROM `urls` WHERE `urls`.`deleted_at` IS NULL AND ((`urls`.`user_id` = ?)) ORDER BY `urls`.`id` LIMIT 1",
		},
		{
			Name:        "Error from DB",
			Query:       "SELECT * FROM `urls` WHERE `urls`.`deleted_at` IS NULL AND ((`urls`.`user_id` = ?)) ORDER BY `urls`.`id` LIMIT 1",
			Err:         errors.New("error"),
			ExpectedErr: errors.New("error"),
		},
	} {
		s.SetupTest()
		s.Run(tt.Name, func() {
			s.Mock.ExpectBegin()
			if tt.Err != nil {
				s.Mock.ExpectQuery(regexp.QuoteMeta(tt.Query)).WillReturnError(tt.Err)
				s.Mock.ExpectRollback()
			} else {
				s.Mock.ExpectQuery(regexp.QuoteMeta(tt.Query)).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "url", "short_url", "user_id"}).AddRow(1, "2021-05-03 15:34:07", "2021-05-03 15:34:07", nil, "https://www.google.com/", "google", 1))
				s.Mock.ExpectCommit()
			}
		})
	}
}

func (s *TestSuiteUrlRepository) TestUpdateUrl() {
	for _, tt := range []struct {
		Name        string
		Query       string
		Err         error
		ExpectedErr error
	}{
		{
			Name:  "Success",
			Query: "UPDATE `urls` SET `updated_at`=?,`url`=?,`short_url`=? WHERE `id` = ?",
		},
		{
			Name:        "Error from DB",
			Query:       "UPDATE `urls` SET `updated_at`=?,`url`=?,`short_url`=? WHERE `id` = ?",
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
	}
}

func (s *TestSuiteUrlRepository) TestDeleteUrl() {
	for _, tt := range []struct {
		Name        string
		Query       string
		Err         error
		ExpectedErr error
	}{
		{
			Name:  "Success",
			Query: "UPDATE `urls` SET `deleted_at`=? WHERE `id` = ?",
		},
		{
			Name:        "Error from DB",
			Query:       "UPDATE `urls` SET `deleted_at`=? WHERE `id` = ?",
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
	}
}

func (s *TestSuiteUrlRepository) TestExpiredUrl() {
	for _, tt := range []struct {
		Name        string
		Query       string
		Err         error
		ExpectedErr error
	}{
		{
			Name:  "Success",
			Query: "SELECT * FROM `urls`.`long_url` WHERE `urls`.`deleted_at` IS NULL AND ((`urls`.`short_url` = ? AND `urls`.`expired_at` > `NOW()`))",
		},
		{
			Name:        "Error from DB",
			Query:       "SELECT * FROM `urls`.`long_url` WHERE `urls`.`deleted_at` IS NULL AND ((`urls`.`short_url` = ? AND `urls`.`expired_at` > `NOW()`))",
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
	}
}

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(TestSuiteUrlRepository))
}
