package routes

import (
	"github.com/labstack/echo/v4"
	"myapp/handlers"
	"myapp/models"
	"myapp/repository"
)

func SetupBooksRoutes(g *echo.Group) {
	repo := repository.NewGormBookRepository(models.DB)
	handler := handlers.NewBookHandler(repo)
	g.GET("/books", handler.GetBooks)
	g.POST("/books", handler.CreateBook)
	g.GET("/books/:id", handler.GetBook)
	g.PUT("/books/:id", handler.UpdateBook)
	g.DELETE("/books/:id", handler.DeleteBook)
}
