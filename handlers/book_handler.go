package handlers

import (
	"github.com/labstack/echo/v4"
	"myapp/dtos"
	"myapp/repository"
	"myapp/utils"
	"net/http"
)

type BookHandler struct {
	Repo repository.IBookRepository
}

func NewBookHandler(repo repository.IBookRepository) *BookHandler {
	return &BookHandler{Repo: repo}
}

func (h *BookHandler) CreateBook(c echo.Context) error {
	book := new(dtos.BookDTO)
	if err := c.Bind(book); err != nil {
		return err
	}

	// TODO: Add validation

	if err := h.Repo.CreateBook(book); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusCreated, book)
}

func (h *BookHandler) DeleteBook(c echo.Context) error {
	strId := c.Param("id")
	id, err := utils.ParseIDToInt(strId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := h.Repo.DeleteBook(id); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *BookHandler) GetBook(c echo.Context) error {
	strId := c.Param("id")
	id, err := utils.ParseIDToInt(strId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	book, err := h.Repo.GetBook(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, book)
}

func (h *BookHandler) GetBooks(c echo.Context) error {
	books, err := h.Repo.GetBooks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, books)
}

func (h *BookHandler) UpdateBook(c echo.Context) error {
	strId := c.Param("id")
	id, err := utils.ParseIDToInt(strId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	book := new(dtos.BookDTO)
	if err := c.Bind(book); err != nil {
		return err
	}

	updatedBook, err := h.Repo.UpdateBook(id, book)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, updatedBook)
}
