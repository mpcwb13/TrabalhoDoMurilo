package service

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"trabalho1/models"
)

var db *gorm.DB

func init() {
	dsn := "host=localhost user=murilo password=123456 dbname=murilo-dev port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Conexão falhou!")
	}
	db.AutoMigrate(&models.Person{})

}

func CreatePerson(person models.Person) (models.Person, error) {
	err := db.Where("email = ?", person.Email).First(&person).Error
	if err == nil {
		return models.Person{}, fmt.Errorf("Ja existe alguem com esse email!")
	}

	err = db.Create(&person).Error
	if err != nil {
		return models.Person{}, err
	}

	return person, nil

}

func UpdatePerson(id int, person models.Person) (*models.Person, error) {

	if person.Email == "" {
		return nil, fmt.Errorf("O email é obrigatório, favor preenchê-lo.")
	}

	if person.Name == "" {
		return nil, fmt.Errorf("O nome é obrigatório, favor preenchê-lo.")
	}

	if person.Age == 0 {
		return nil, fmt.Errorf("A idade é obrigatória, favor preenchê-la.")
	}

	var dbPerson models.Person //variavel banco de dados
	err := db.Where("id = ?", id).First(&dbPerson).Error
	if err != nil {
		return nil, fmt.Errorf("Não existe ninguém com esse id cadastrado")
	}

	dbPerson.Name = person.Name
	dbPerson.Age = person.Age
	dbPerson.Email = person.Email
	dbPerson.Birthday = person.Birthday
	dbPerson.Gender = person.Gender

	err = db.Save(&dbPerson).Error
	if err != nil {
		return nil, fmt.Errorf("Erro ao salvar a pessoa no banco de dados")
	}

	return nil, err
}

func DeletedPeople(id int, person models.Person) (*models.Person, error) {
	err := db.Where("id = ?", id).First(&person).Error
	if err != nil {
		return nil, fmt.Errorf("Não existe ninguem com esse id cadastrado")
	}

	err = db.Delete(&person).Error
	if err != nil {
		return nil, err
	}

	return nil, err
}

func SearchPeople(id int) (*models.Person, error) {
	var person models.Person
	err := db.Where("id =  ?", id).First(&person).Error
	if err != nil {
		return nil, fmt.Errorf("Não existe ninguem com esse id cadastrado")
	}

	return &person, nil

}

func SearchAllPeople() ([]models.Person, error) {
	var people []models.Person
	err := db.Find(&people).Error
	if err != nil {
		return nil, err
	}

	if len(people) == 0 {

		return nil, fmt.Errorf("Nenhuma pessoa encontrada")
	}
	return people, nil
}
