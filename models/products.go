package models

import "gorm.io/gorm"

type Products struct {
	gorm.Model

	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
}
