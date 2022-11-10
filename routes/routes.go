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

	// Non-Middleware
	e.POST("/register", userController.Create())
	e.POST("/login", userController.Auth())
	e.GET("/:short", urlController.FindUrl())

	// Load ENV
	data := config.LoadENV()

	// Group User
	user := e.Group("/users")
	user.Use(middleware.JWT([]byte(data["jwtSecret"])))
	user.GET("", urlController.UserUrl())

	// Group Url
	url := e.Group("/urls")
	url.Use(middleware.JWT([]byte(data["jwtSecret"])))
	url.POST("", urlController.GenerateUrl())
	url.GET("", urlController.GetAllUrl())
	url.PUT("/:short", urlController.UpdateUrl())
	url.DELETE("/:short", urlController.DeleteUrl())

	return e

}
