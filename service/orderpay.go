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
	db.AutoMigrate(&models.OrderPay{})

}

func CreateOrder(order models.OrderPay) (models.OrderPay, error) {
	err := db.Create(&order).Error
	if err != nil {
		return models.OrderPay{}, err
	}
	return order, nil
}
