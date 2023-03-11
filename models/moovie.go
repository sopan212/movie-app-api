package models

import "gorm.io/gorm"

type Moovie struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Poster      string  `json:"poster"`
	Rating      float32 `json:"rating"`
	CategorieID uint    `json:"category_id"`
	Categorie   Categorie
}
