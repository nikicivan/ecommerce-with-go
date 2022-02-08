package controllers

import (
	"ecommerce/src/database"
	"ecommerce/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetLinksByUserId returns all links by user id
func GetLinksByUserId(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var links []models.Link

	database.DB.Where("user_id = ?", id).Find(&links)

	for i, link := range links {
		var orders []models.Order
		database.DB.Where("code = ? and complete = true", link.Code).Find(&orders)

		links[i].Orders = orders
	}

	return c.JSON(links)
}
