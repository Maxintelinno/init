package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	e.GET("/*", func(c echo.Context) error {
		return c.String(http.StatusOK, "API Running")
	})

	// 👇 สำคัญมาก
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start("0.0.0.0:" + port))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
