package controller

import (
	"GO_MINIPROJECT/repository"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type Logincontroller struct {
	iGymrepo repository.IGymRepo
}

func NewLoginController(iGymrepo repository.IGymRepo) Logincontroller {
	return Logincontroller{iGymrepo}
}

func (ctrl Logincontroller) Login(c echo.Context) error {
	authorization := c.Request().Header.Get("authorization")
	arrAuth := strings.Split(authorization, " ")
	if len(arrAuth) != 2 {
		fmt.Println("header auth tidak valid")
		return c.JSON(400, map[string]interface{}{
			"message": "header auth tidak valid",
		})
	} else if arrAuth[0] != "Basic" {
		fmt.Println("header auth must be basic")
		return c.JSON(400, map[string]interface{}{
			"message": "header auth must be basic",
		})
	}

	var decodeByte, _ = base64.StdEncoding.DecodeString(arrAuth[1])
	arrDecodeString := strings.Split(string(decodeByte), ":")
	username, password := arrDecodeString[0], arrDecodeString[1]
	_, err := ctrl.iGymrepo.GetUserUsernamePassword(username, password)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "email / password not registered",
		})
	}
	jwtToken := ctrl.CreateJWT(username)
	return c.JSON(200, map[string]interface{}{
		"token": jwtToken,
	})
}

func (ctrl Logincontroller) CreateJWT(username string) string {
	secret := []byte("secret123")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username":   username,
			"createdAt":  time.Now(),
			"expired_at": time.Now().Add(1 * time.Hour),
		})
	tokenString, err := token.SignedString(secret)
	fmt.Println("2")
	if err != nil {
		fmt.Println(err)
	}
	return tokenString
}
