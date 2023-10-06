package models

import (
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model

	Name     string      `json:"name"`
	Age      int         `json:"age"`
	Email    string      `json:"email"`
	Birthday string      `json:"birthday"`
	Gender   string      `json:"gender"`
	OrderPay []*OrderPay `gorm:"foreignKey:PersonID"`
}
