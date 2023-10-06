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

func CreateProduct(products models.Products) (models.Products, error) {

	err := db.Create(&products).Error
	if err != nil {
		return models.Products{}, err
	}

	return products, nil

}

func UpdateProduct(id int, product models.Products) (*models.Products, error) {

	if product.Name == "" {
		return nil, fmt.Errorf("O nome do produto é obrigatorio")
	}

	if product.Price == 0 {
		return nil, fmt.Errorf("O preço do produto é obrigatorio")
	}

	if product.Description == "" {
		return nil, fmt.Errorf("A descrição do produto é obrigatoria")
	}

	if product.Quantity == 0 {
		return nil, fmt.Errorf("A quantidade é obrigatoria")
	}

	var dbProduct models.Products
	err := db.Where("id = ?", id).First(&dbProduct).Error
	if err != nil {
		return nil, fmt.Errorf("Produto não encontado.")
	}

	dbProduct.Name = product.Name
	dbProduct.Price = product.Price
	dbProduct.Description = product.Description
	dbProduct.Quantity = product.Quantity

	err = db.Save(&dbProduct).Error
	if err != nil {
		return nil, fmt.Errorf("Erro na hora de salvar no banco.")
	}

	return nil, err

}

func SearchProduct(id int) (*models.Products, error) {
	var dbProduct models.Products
	err := db.Where("id = ?", id).First(&dbProduct).Error
	if err != nil {
		return nil, fmt.Errorf("Produto não encontrado")
	}

	return &dbProduct, nil

}

func SearchAllProducts() ([]models.Products, error) {
	var dbProducts []models.Products
	err := db.Find(&dbProducts).Error
	if err != nil {
		return nil, err
	}

	if len(dbProducts) == 0 {
		return nil, fmt.Errorf("Nenhum produto encontrado")
	}

	return dbProducts, nil

}

func DeletedProduct(id int, products models.Products) (*models.Products, error) {

	err := db.Where("id = ?", id).First(&products).Error
	if err != nil {
		return nil, fmt.Errorf("Não existe ninguem com esse id")
	}

	err = db.Delete(&products).Error
	if err != nil {
		return nil, err
	}

	return nil, err

}
