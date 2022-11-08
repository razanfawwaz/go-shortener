package controller

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"urlshortener/domain"
	"urlshortener/helper"
	"urlshortener/url/usecase"
)

type UrlController interface{}

type urlController struct {
	urlUsecase usecase.UseCase
}

func NewUrlController(urlUsecase usecase.UseCase) *urlController {
	return &urlController{urlUsecase}
}

func (u *urlController) GenerateUrl(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var url domain.Url
		id := helper.ClaimToken(c)
		url.UserID = int(id)
		if err := c.Bind(&url); err != nil {
			return c.JSON(400, echo.Map{
				"message": err.Error(),
			})
		}

		url, err := u.urlUsecase.GenerateUrl(url)
		{
			if err != nil {
				return c.JSON(500, echo.Map{
					"message": err.Error(),
				})
			}
		}
		return c.JSON(200, echo.Map{
			"message": "success",
			"data":    url,
		})
	}
}

func (u *urlController) GetAllUrl(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		urls, err := u.urlUsecase.GetAllUrl()
		{
			if err != nil {
				return c.JSON(500, echo.Map{
					"message": err.Error(),
				})
			}
		}

		return c.JSON(200, echo.Map{
			"message": "success",
			"data":    urls,
		})
	}
}

func (u *urlController) FindUrl(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		short := c.Param("short")
		url, err := u.urlUsecase.FindUrl(short)
		{
			if err != nil {
				return c.JSON(500, echo.Map{
					"message": err.Error(),
				})
			}
		}
		return c.Redirect(301, url)
	}
}
func (u *urlController) UpdateUrl(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		short := c.Param("short")

		var url domain.Url
		id := helper.ClaimToken(c)
		url.UserID = int(id)

		if err := c.Bind(&url); err != nil {
			return c.JSON(400, echo.Map{
				"message": err.Error(),
			})
		}

		url, err := u.urlUsecase.UpdateUrl(short, int(id), url)
		{
			if err != nil {
				return c.JSON(500, echo.Map{
					"message": err.Error(),
				})
			}
		}

		return c.JSON(200, echo.Map{
			"message": "success",
			"data":    url,
		})
	}
}

func (u *urlController) DeleteUrl(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var url domain.Url
		if err := c.Bind(&url); err != nil {
			return c.JSON(400, echo.Map{
				"message": err.Error(),
			})
		}

		url, err := u.urlUsecase.DeleteUrl(url)
		{
			if err != nil {
				return c.JSON(500, echo.Map{
					"message": err.Error(),
				})
			}
		}

		return c.JSON(200, echo.Map{
			"message": "success",
			"data":    url,
		})
	}
}

func (u *urlController) UserUrl(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := helper.ClaimToken(c)
		urls, err := u.urlUsecase.UserUrl(int(id))
		{
			if err != nil {
				return c.JSON(500, echo.Map{
					"message": err.Error(),
				})
			}
		}

		return c.JSON(200, echo.Map{
			"message": "success",
			"data":    urls,
		})
	}
}
