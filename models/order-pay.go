package models

import "gorm.io/gorm"

type OrderPay struct {
	gorm.Model

	NumberOrder int         `json:"numberOrder"`
	StatusOrder string      `json:"statusOrder"`
	CancelOrder bool        `json:"cancelOrder"`
	Total       float64     `json:"total" gorm:"-"`
	PersonID    int         `json:"personID" gorm:"-"`
	Products    []*Products `gorm:"many2many:orderpay_products;"`
}
