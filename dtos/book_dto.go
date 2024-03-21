package dtos

type BookDTO struct {
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Rating      float64 `json:"rating"`
	IsPublished bool    `json:"isPublished"`
}
