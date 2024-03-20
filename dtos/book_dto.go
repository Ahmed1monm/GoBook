package dtos

type BookDTO struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Rating      int    `json:"rating"`
	IsPublished bool   `json:"is-published"`
}
