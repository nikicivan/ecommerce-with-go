package controllers

import (
	"context"
	"ecommerce/src/database"
	"ecommerce/src/models"
	"encoding/json"
	"sort"
	"strconv"
	"strings"
	"time"

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

	// delete cahce from redis with gorutines and channels
	go database.ClearCache("products_backoffice", "products_frontend")

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

// // DeleteCache deletes cache in redis
// func DeleteCache(key string) {
// 	database.Cache.Del(context.Background(), key)
// }

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

	// // delete cache with gorutines
	// go DeleteCache("products_backoffice")
	// go DeleteCache("products_frontend")

	// delete cahce from redis with gorutines and channels
	go database.ClearCache("products_backoffice", "products_frontend")

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

	// delete cahce from redis with gorutines and channels
	go database.ClearCache("products_backoffice", "products_frontend")

	c.SendStatus(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

// ProductFrontend returns all products for frontend with redis caching
func ProductFrontend(c *fiber.Ctx) error {
	var products []models.Product
	var ctx = context.Background()

	result, err := database.Cache.Get(ctx, "products_frontend").Result()

	if err != nil {
		database.DB.Find(&products)

		bytes, err := json.Marshal(products)

		if err != nil {
			return err
		}

		database.Cache.Set(ctx, "products_frontend", bytes, 30*time.Minute)

	} else {
		json.Unmarshal([]byte(result), &products)
	}

	return c.JSON(products)
}

// ProductBackoffice returns all products for backoffice with redis caching
func ProductBackoffice(c *fiber.Ctx) error {
	var products []models.Product
	var ctx = context.Background()

	result, err := database.Cache.Get(ctx, "products_backoffice").Result()

	if err != nil {
		database.DB.Find(&products)

		bytes, err := json.Marshal(products)

		if err != nil {
			return err
		}

		database.Cache.Set(ctx, "products_backoffice", bytes, 30*time.Minute)

	} else {
		json.Unmarshal([]byte(result), &products)
	}

	var searchedProducts []models.Product

	if s := c.Query("s"); s != "" {
		lower := strings.ToLower(s)

		for _, product := range products {
			if strings.Contains(strings.ToLower(product.Title), lower) || strings.Contains(strings.ToLower(product.Description), lower) {
				searchedProducts = append(searchedProducts, product)
			}
		}
	} else {
		searchedProducts = products
	}

	if sortParam := c.Query("sort"); sortParam != "" {
		sortLower := strings.ToLower(sortParam)

		if sortLower == "asc" {
			sort.Slice(searchedProducts, func(i, j int) bool {
				return searchedProducts[i].Price < searchedProducts[j].Price
			})
		}

		if sortLower == "desc" {
			sort.Slice(searchedProducts, func(i, j int) bool {
				return searchedProducts[i].Price > searchedProducts[j].Price
			})
		}
	}

	var total = len(searchedProducts)
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage := 9

	var data []models.Product

	if total <= page*perPage && total > (page-1)*perPage {
		data = searchedProducts[(page-1)*perPage : total]
	} else if total >= page*perPage {
		data = searchedProducts[(page-1)*perPage : page*perPage]
	} else {
		data = []models.Product{}
	}

	return c.JSON(fiber.Map{
		"data":      data,
		"total":     total,
		"page":      page,
		"last_page": total/perPage + 1,
	})
}
