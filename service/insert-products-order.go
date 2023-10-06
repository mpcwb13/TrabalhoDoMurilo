package service

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"trabalho1/models"
)

func init() {
	dsn := "host=localhost user=murilo password=123456 dbname=murilo-dev port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Conexão falhou!")
	}
	db.AutoMigrate(&models.Products{})

}

func InsertProducts(orderId int, productsIds []int) (*models.OrderPay, error) {
	var dbOrder models.OrderPay
	err := db.Where("number_order = ?", orderId).Find(&dbOrder).Error
	if err != nil {
		return nil, fmt.Errorf("Não conseguimos localizar a ordem")
	}

	var dbProducts []models.Products
	err = db.Where("id IN ?", productsIds).Find(&dbProducts).Error
	if err != nil {
		return nil, fmt.Errorf("Não conseguimos localizar o produto")
	}

	err = db.Model(&dbOrder).Association("Products").Append(&dbProducts)
	if err != nil {
		return nil, err
	}

	err = db.Preload("Products").Find(&dbOrder, orderId).Error
	if err != nil {
		return nil, fmt.Errorf("Não conseguimos localizar a ordem")
	}
	var total float64
	for _, product := range dbOrder.Products {
		total += product.Price
	}

	dbOrder.Total = total

	return &dbOrder, err
}
