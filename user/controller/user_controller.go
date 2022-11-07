package controller

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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

func (u *userController) Create(db *gorm.DB) echo.HandlerFunc {
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
			"message": "success",
			"data":    user,
		})
	}
}

func (u *userController) Auth(db *gorm.DB) echo.HandlerFunc {
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
		token, err := helper.CreateToken(user.ID, user.Email, c)
		if err != nil {
			return c.JSON(500, echo.Map{
				"message": err.Error(),
			})
		}

		c.Set("userID", user.ID)
		c.Set("userEmail", user.Email)

		return c.JSON(200, echo.Map{
			"message": "success",
			"data":    user,
			"token":   token,
		})
	}
}
