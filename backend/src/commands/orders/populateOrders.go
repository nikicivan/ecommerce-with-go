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
		var orderItems []models.OrderItem

		for j := 0; j < rand.Intn(5); j++ {
			price := float64(rand.Intn(90) + 10)
			quantity := uint(rand.Intn(5))

			orderItems = append(orderItems, models.OrderItem{
				ProductTitle:      faker.Word(),
				Price:             price,
				Quantity:          quantity,
				AdminRevenue:      0.9 * price * float64(quantity),
				AmbassadorRevenue: 0.1 * price * float64(quantity),
			})
		}

		database.DB.Create(&models.Order{
			UserId:          uint(rand.Intn(60) + 1),
			Code:            faker.Username(),
			AmbassadorEmail: faker.Email(),
			FirstName:       faker.FirstName(),
			LastName:        faker.LastName(),
			Email:           faker.Email(),
			Complete:        true,
			OrderItems:      orderItems,
		})

	}
	log.Println("Fake orders have been populated....")
}
