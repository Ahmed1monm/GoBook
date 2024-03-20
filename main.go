package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"myapp/models"
	"myapp/routes"
)

func main() {
	e := echo.New()

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// db connection
	db := models.InitDB("books.db")
	models.Migrate(db)
	// Routes
	api := e.Group("/api")
	routes.SetupBooksRoutes(api)
	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
