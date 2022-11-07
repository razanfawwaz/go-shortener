package helper

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"time"
	"urlshortener/config"
)

func CreateToken(userID uint, userEmail string, ctx echo.Context) (string, error) {
	claims := jwt.MapClaims{
		"authorization": true,
		"userID":        userID,
		"userEmail":     userEmail,
		"exp":           time.Now().Add(time.Minute * 30).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	data := config.LoadENV()
	secret := data["jwtSecret"]

	return token.SignedString([]byte(secret))
}
