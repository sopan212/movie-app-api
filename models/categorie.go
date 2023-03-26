package models

import (
	"time"

	"gorm.io/gorm"
)

type Categorie struct {
	gorm.Model
	CategoryName string           `json:"category_name"`
	MoovieID     uint             `json:"moovie_id"`
	Moovies      []CategoryMoovie `json:"moovies"`
}

type ResponseCategory struct {
	ID           uint   `json:"id"`
	Categoryname string `json:"category_name"`
}
type ResponseForCategory struct {
	ID           uint   `json:"id"`
	Categoryname string `json:"category_name"`
	Moovies      []Moovie
}
type CategoryMoovie struct {
	gorm.Model
	CategoryID uint      `json:"category_id"`
	Categorie  Categorie `gorm:"foreignKey:CategoryID;reference:ID"`
	MoovieID   uint      `json:"moovie_id"`
	Moovie     Moovie    `gorm:"foreignKey:MoovieID;reference:ID"`
	Decription string    `json:"description"`
}

type ResponseCategoryMoovie struct {
	Id           uint   `json:"id"`
	CategoryName string `json:"category_name"`
	MoovieName   string `json:"moovie_name"`
	Description  string `json:"description"`
	CreatedAt    time.Time
}
