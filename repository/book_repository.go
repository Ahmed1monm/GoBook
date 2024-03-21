package repository

import (
	"myapp/dtos"
	"myapp/models"
)

type IBookRepository interface {
	CreateBook(book *dtos.BookDTO) error
	DeleteBook(id uint) error
	GetBook(id uint) (*models.Book, error)
	GetBooks() ([]models.Book, error)
	UpdateBook(id uint, dto *dtos.BookDTO) (*models.Book, error)
}
