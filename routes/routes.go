package routes

import (
	"github.com/labstack/echo/v4"
	"urlshortener/config"
	"urlshortener/url/controller"
	"urlshortener/url/repository"
	"urlshortener/url/usecase"
)

func New() *echo.Echo {
	e := echo.New()
	urlRepo := repository.NewUrlRepository(config.DB)
	urlService := usecase.NewUrlUsecase(urlRepo)
	urlController := controller.NewUrlController(urlService)

	e.POST("/urls", urlController.GenerateUrl(config.DB))
	// get by using params
	e.GET("/urls/:short", urlController.FindUrl(config.DB))
	e.GET("/urls", urlController.GetAllUrl(config.DB))
	e.PUT("/urls/:short", urlController.UpdateUrl(config.DB))
	return e
}
