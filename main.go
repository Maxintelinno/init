package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	log.Println("ENV CHECK:", os.Getenv("DATABASE_URL"))

	// 🔥 Connect DB
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("✅ Connected to PostgreSQL")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "API Running")
	})

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT not set")
	}

	e.Logger.Fatal(e.Start("0.0.0.0:" + port))
}
