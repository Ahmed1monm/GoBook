package handlers

import (
	"github.com/labstack/echo/v4"
	"myapp/repository"
)

type BookHandler struct {
	Repo repository.IBookRepository
}

func NewBookHandler(repo repository.IBookRepository) *BookHandler {
	return &BookHandler{Repo: repo}
}

func (h *BookHandler) CreateBook(c echo.Context) error {
	return error(nil)
}

func (h *BookHandler) DeleteBook(c echo.Context) error {
	return error(nil)
}

func (h *BookHandler) GetBook(c echo.Context) error {
	return error(nil)
}

func (h *BookHandler) GetBooks(c echo.Context) error {
	return error(nil)
}

func (h *BookHandler) UpdateBook(c echo.Context) error {
	return error(nil)
}
