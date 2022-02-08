package main

import (
	"ecommerce/src/database"
	"ecommerce/src/models"
	"log"

	"github.com/bxcodec/faker/v3"
)

func main() {
	database.Connect()

	for i := 0; i < 30; i++ {
		ambasador := models.User{
			FirstName:   faker.FirstName(),
			LastName:    faker.LastName(),
			Email:       faker.Email(),
			IsAmbasador: true,
		}

		ambasador.SetPassword("123456")

		database.DB.Create(&ambasador)
	}
	log.Println("Fake ambasador users have been populated....")
}
