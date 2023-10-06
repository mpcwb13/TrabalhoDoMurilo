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
	db.AutoMigrate(&models.Person{})

}

func DeleteProductOrder(orderId int, productId int) (*models.OrderPay, error) {
	var dbOrder models.OrderPay
	err := db.Where("number_order = ?", orderId).Find(&dbOrder).Error
	if err != nil {
		return nil, fmt.Errorf("Não conseguimos localizar a ordem")
	}

	var dbProduct models.Products
	err = db.Where("id = ?", productId).Find(&dbProduct).Error
	if err != nil {
		return nil, fmt.Errorf("Não conseguimos localizar o produto")
	}

	err = db.Model(&dbOrder).Association("Products").Delete(dbProduct)
	if err != nil {
		return nil, fmt.Errorf("Não conseguimos deletar o produto")
	}

	return &dbOrder, err
}
