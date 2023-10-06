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

func ListOrderOpen() ([]models.OrderPay, error) {
	var dbStatus []models.OrderPay
	err := db.Where("status_order = ?", dbStatus).First(&dbStatus).Error
	if err != nil {
		return nil, err
	}

	return dbStatus, nil
}
