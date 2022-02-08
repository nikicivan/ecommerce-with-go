package controllers

import (
	"ecommerce/src/database"
	"ecommerce/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetProducts returns all products
func GetProducts(c *fiber.Ctx) error {
	var products []models.Product

	database.DB.Find(&products)

	return c.JSON(products)
}

// CreateProduct creates a new product in db
func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	err := c.BodyParser(&product)
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "You did not provide all required fields",
		})
	}

	database.DB.Create(&product)
	c.SendStatus(fiber.StatusCreated)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

// GetProduct returns product by ID
func GetProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var product models.Product

	product.Id = uint(id)

	database.DB.Find(&product)

	return c.JSON(product)
}

// UpdateProduct updates product by ID
func UpdateProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var product models.Product

	product.Id = uint(id)

	err := c.BodyParser(&product)
	if err != nil {
		return err
	}

	database.DB.Model(&product).Where("id = ?", product.Id).Updates(&product)

	c.SendStatus(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

// DeleteProduct removes product by ID from DB
func DeleteProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var product models.Product
	product.Id = uint(id)

	database.DB.Model(&product).Where("id = ?", product.Id).Delete(&product)

	c.SendStatus(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}
