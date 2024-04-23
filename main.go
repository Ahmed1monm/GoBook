package main

import (
	"fmt"
	"myapp/models"
	"myapp/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	fmt.Fprint(e.Logger.Output(), "Hello, World!")
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
