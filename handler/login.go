package handler

import (
	"log"
	"net/http"
	"time"
	"github.com/dgrijalva/jwt-go"
	"example.com/m/models"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	// req := new(models.LoginRequest)
	// c.Bind(req) //bóc tách dữ liệu trong request
	// log.Printf("req data %+v", req)

	username := c.Get("username").(string)
	log.Printf("login with username: %s", username)

	admin := c.Get("admin").(bool)
	log.Printf("login with admin: %t", admin)

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["admin"] = admin
	claims["exp"] = time.Now().Add(2 * time.Hour).Unix()

	t, err := token.SignedString([]byte("mykey"))

	if err != nil {
		log.Printf("error token signing", err)
		return err
	}

	return c.JSON(http.StatusOK, &models.LoginResponse{
		Token: t,
	})
}
