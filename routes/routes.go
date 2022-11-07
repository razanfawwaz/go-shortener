package routes

import (
	"github.com/labstack/echo/v4"
	"urlshortener/config"
	urlCon "urlshortener/url/controller"
	urlRep "urlshortener/url/repository"
	urlUsc "urlshortener/url/usecase"
	userCon "urlshortener/user/controller"
	userRep "urlshortener/user/repository"
	userUsc "urlshortener/user/usecase"
)

func New() *echo.Echo {
	e := echo.New()
	urlRepo := urlRep.NewUrlRepository(config.DB)
	urlService := urlUsc.NewUrlUsecase(urlRepo)
	urlController := urlCon.NewUrlController(urlService)

	userRepo := userRep.NewUserRepository(config.DB)
	userService := userUsc.NewUserUsecase(userRepo)
	userController := userCon.NewUserController(userService)

	e.POST("/urls", urlController.GenerateUrl(config.DB))
	e.GET("/urls/:short", urlController.FindUrl(config.DB))
	e.GET("/urls", urlController.GetAllUrl(config.DB))
	e.PUT("/urls/:short", urlController.UpdateUrl(config.DB))

	e.POST("/users", userController.Create(config.DB))
	e.POST("/users/login", userController.Auth(config.DB))

	return e
}
