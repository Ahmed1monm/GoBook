package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Rating      float64 `json:"rating"`
	IsPublished bool    `json:"is-published"`
}
