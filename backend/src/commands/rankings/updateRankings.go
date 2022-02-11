package main

import (
	"context"
	"ecommerce/src/database"
	"ecommerce/src/models"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	database.Connect()
	database.SetupRedis()

	ctx := context.Background()

	var users []models.User

	database.DB.Find(&users, models.User{
		IsAmbasador: true,
	})

	for _, user := range users {
		ambassador := models.Ambassador(user)

		ambassador.CalculateRevenue(database.DB)

		database.Cache.ZAdd(ctx, "rankings", &redis.Z{
			Score:  *ambassador.Revenue,
			Member: user.FullName(),
		})

		fmt.Println("Caching rankings in redis have been done...")
	}

}
