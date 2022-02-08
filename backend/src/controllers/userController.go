package controllers

import (
	"ecommerce/src/database"
	"ecommerce/src/models"

	"github.com/gofiber/fiber/v2"
)

// GetAmbasadors returns all ambasador users
func GetAmbasadors(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Where("is_ambasador = true").Find(&users)

	return c.JSON(users)
}
