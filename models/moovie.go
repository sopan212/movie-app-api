package models

import (
	"gorm.io/gorm"
)

type Moovie struct {
	gorm.Model
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Years       int              `json:"years"`
	Poster      string           `json:"poster"`
	CategorieID uint             `json:"category_id"`
	Categories  []CategoryMoovie `json:"categories"`
	Reviews     []Review
	Rating      []Ratting `json:"rating"`
}

type ResponseMoovies struct {
	ID          uint               `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Years       int                `json:"years"`
	Poster      string             `json:"poster"`
	Rating      []ResponseRatting  `json:"rating"`
	Categories  []ResponseCategory `json:"categories"`
	Reviews     []ResponseReview
}

//response untuk category
type ResponseMooviesForCategory struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Years       int    `json:"years"`
	Poster      string `json:"poster"`
}
