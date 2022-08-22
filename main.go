package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/labstack/echo/v4"
	"github.com/ralberts/golang-starter-template/database"
	"github.com/ralberts/golang-starter-template/examples"
)

func main() {
	log.SetOutput(os.Stdout)

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	dburl := os.Getenv("DATABASE_URL")

	if dburl == "" {
		log.Fatal("$DATABASE_URL must be set")
	}

	sslmode := os.Getenv("SSL_MODE")

	if sslmode == "" {
		// options here: https://godoc.org/github.com/lib/pq
		sslmode = "require"
	}

	// testStr := database.Test()
	// database.Connect(dburl)
	db := database.Connect(dburl, sslmode)

	e := echo.New()

	group := "/api"

	// example migration
	example.Migrate(db)
	example.Register(group, e)

	log.Print("Done registering routes")
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))

	// v1 := r.Group("/api")

	// auth.Register(v1.Group("/auth"))
}
