package controller

import (
	"github.com/labstack/echo/v4"
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

func (u *urlController) GenerateUrl() echo.HandlerFunc {
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
			"data": url,
		})
	}
}

func (u *urlController) GetAllUrl() echo.HandlerFunc {
	return func(c echo.Context) error {

		urls, err := u.urlUsecase.GetAllUrl()
		{
			if err != nil {
				return c.JSON(400, echo.Map{
					"message": err.Error(),
				})
			}
		}

		return c.JSON(200, echo.Map{
			"data": urls,
		})
	}
}

func (u *urlController) FindUrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		short := c.Param("short")
		url, err := u.urlUsecase.FindUrl(short)
		{
			if err != nil {
				return c.JSON(400, echo.Map{
					"message": err.Error(),
				})
			}
		}

		return c.Redirect(301, url)
	}
}
func (u *urlController) UpdateUrl() echo.HandlerFunc {
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
			"data": url,
		})
	}
}

func (u *urlController) DeleteUrl() echo.HandlerFunc {
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

		err := u.urlUsecase.DeleteUrl(short, int(id), url)
		{
			if err != nil {
				return c.JSON(500, echo.Map{
					"message": err.Error(),
				})
			}
		}

		return c.JSON(200, echo.Map{
			"message": "Url deleted",
		})
	}
}

func (u *urlController) UserUrl() echo.HandlerFunc {
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
			"data": urls,
		})
	}
}
