package repository

type IBookRepository interface {
	CreateBook() error
	DeleteBook() error
	GetBook() error
	GetBooks() error
	UpdateBook() error
}
