package repository

import (
	"gorm.io/gorm"
	"myapp/dtos"
	"myapp/models"
)

type GormBookRepository struct {
	DB *gorm.DB
}

func NewGormBookRepository(db *gorm.DB) *GormBookRepository {
	return &GormBookRepository{DB: db}
}

func (r *GormBookRepository) CreateBook(book *dtos.BookDTO) error {
	b := models.Book{}
	b.Title = book.Title
	b.Author = book.Author
	b.IsPublished = book.IsPublished
	return r.DB.Create(&b).Error
}

func (r *GormBookRepository) DeleteBook(id uint) error {
	return r.DB.Delete(&models.Book{}, id).Error
}

func (r *GormBookRepository) GetBook(id uint) (*models.Book, error) {
	var book models.Book
	err := r.DB.First(&book, id).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *GormBookRepository) GetBooks() ([]models.Book, error) {
	var books []models.Book
	err := r.DB.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *GormBookRepository) UpdateBook(id uint, book *dtos.BookDTO) (*models.Book, error) {
	var b *models.Book
	err := r.DB.First(&b, id).Error
	if err != nil {
		return b, err
	}
	b.Title = book.Title
	b.Author = book.Author
	b.IsPublished = book.IsPublished
	err = r.DB.Save(&b).Error
	if err != nil {
		return b, err
	}
	return b, nil
}
