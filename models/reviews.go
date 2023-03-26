package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	Review   string `json:"review"`
	MoovieID uint   `json:"moovie_id"`
	Moovie   Moovie
}

type ResponseReview struct {
	ID     uint   `json:"id"`
	Review string `json:"review"`
}
