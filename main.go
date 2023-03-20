package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"example.com/m/handler"
	"example.com/m/mdw"
)


func main() {
	server := echo.New()
	server.Use(middleware.Logger()) // log trên server về các request

	isLoggedIn := middleware.JWT([]byte("mykey"))

	server.GET("/", handler.Hello, isLoggedIn)

	server.POST("/login", handler.Login, middleware.BasicAuth(mdw.BasicAuth))
	
	server.Logger.Fatal(server.Start(":8888"))
}



