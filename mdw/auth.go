package mdw

import (
	"github.com/labstack/echo/v4"
)

func BasicAuth(username string, password string, c echo.Context) (bool, error){
	if username == "admin" || password == "password" {
		c.Set("username", username)
		c.Set("admin", true)
		c.Set("username", username)
		return true, nil
	}
	if username == "khai" || password == "password" {
		c.Set("username", username)
		c.Set("admin", false)
		c.Set("username", username)
		return true, nil
	}
	return false, nil
}
