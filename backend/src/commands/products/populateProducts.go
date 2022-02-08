package main

import (
	"ecommerce/src/database"
	"ecommerce/src/models"
	"log"
	"math/rand"

	"github.com/bxcodec/faker/v3"
)

func main() {
	database.Connect()

	for i := 0; i < 30; i++ {
		product := models.Product{
			Title:       faker.Username(),
			Description: faker.Username(),
			Image:       faker.URL(),
			Price:       float64(rand.Intn(90) + 10),
		}

		database.DB.Create(&product)

	}
	log.Println("Fake products have been populated....")
}
