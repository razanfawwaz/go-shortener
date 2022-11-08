package routes

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"urlshortener/config"
	urlCon "urlshortener/url/controller"
	urlRep "urlshortener/url/repository"
	urlUsc "urlshortener/url/usecase"
	userCon "urlshortener/user/controller"
	userRep "urlshortener/user/repository"
	userUsc "urlshortener/user/usecase"
)

type UserValidator struct {
	validator *validator.Validate
}

func (v *UserValidator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func New() *echo.Echo {
	e := echo.New()
	e.Validator = &UserValidator{validator: validator.New()}

	// URL
	urlRepo := urlRep.NewUrlRepository(config.DB)
	urlService := urlUsc.NewUrlUsecase(urlRepo)
	urlController := urlCon.NewUrlController(urlService)

	// User
	userRepo := userRep.NewUserRepository(config.DB)
	userService := userUsc.NewUserUsecase(userRepo)
	userController := userCon.NewUserController(userService)

	e.GET("/:short", urlController.FindUrl(config.DB))
	e.POST("/users", userController.Create(config.DB))
	e.POST("/users/login", userController.Auth(config.DB))

	auth := e.Group("")
	data := config.LoadENV()
	auth.Use(middleware.JWT([]byte(data["jwtSecret"])))
	auth.GET("/users", urlController.UserUrl(config.DB))
	auth.POST("/urls", urlController.GenerateUrl(config.DB))
	auth.GET("/urls", urlController.GetAllUrl(config.DB))
	auth.PUT("/urls/:short", urlController.UpdateUrl(config.DB))

	return e

}
