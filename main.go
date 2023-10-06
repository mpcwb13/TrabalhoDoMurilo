package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"trabalho1/models"
)

func main() {
	dsn := "host=localhost user=murilo password=123456 dbname=murilo-dev port=5432 sslmode=disable TimeZone=America/Sao_Paulo"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Conexão falhou!")
	}

	db.AutoMigrate(&models.Person{})

	db.Create(&models.Person{Name: "Murilo", Age: 34, Email: "murilo@email.com"})

	var person models.Person
	db.First(&person, 1)
	db.First(&person, "name = ? ", "Murilo")

}
