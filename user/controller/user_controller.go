package controller

import (
	"github.com/labstack/echo/v4"
	"urlshortener/domain"
	"urlshortener/helper"
	"urlshortener/user/usecase"
)

type UserController interface{}

type userController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) *userController {
	return &userController{userUsecase}
}

func (u *userController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var user domain.User
		if err := c.Bind(&user); err != nil {
			return c.JSON(400, echo.Map{
				"message": err.Error(),
			})
		}

		user, err := u.userUsecase.Create(user)
		{
			if err != nil {
				return c.JSON(500, echo.Map{
					"message": err.Error(),
				})
			}
		}

		return c.JSON(200, echo.Map{
			"data": user,
		})
	}
}

func (u *userController) Auth() echo.HandlerFunc {
	return func(c echo.Context) error {
		var user domain.User
		if err := c.Bind(&user); err != nil {
			return c.JSON(400, echo.Map{
				"message": err.Error(),
			})
		}

		user, err := u.userUsecase.Auth(user)
		{
			if err != nil {
				return c.JSON(500, echo.Map{
					"message": err.Error(),
				})
			}
		}
		// create token jwt
		token, err := helper.CreateToken(user.ID)
		if err != nil {
			return c.JSON(500, echo.Map{
				"message": err.Error(),
			})
		}

		c.Set("userID", user.ID)
		c.Set("userEmail", user.Email)

		return c.JSON(200, echo.Map{

			"data":  user,
			"token": token,
		})
	}
}
