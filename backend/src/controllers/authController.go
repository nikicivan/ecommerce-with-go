package controllers

import (
	"ecommerce/src/database"
	"ecommerce/src/middlewares"
	"ecommerce/src/models"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Register creates new user
func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}

	user := models.User{
		FirstName:   data["first_name"],
		LastName:    data["last_name"],
		Email:       data["email"],
		IsAmbasador: strings.Contains(c.Path(), "/api/ambassador"),
	}

	user.SetPassword(data["password"])

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id > 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "User already exist",
		})
	}

	database.DB.Create(&user)

	return c.JSON(user)
}

// Login sing in existing user
func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "User does not exist",
		})
	}

	err := user.CompareHashedPassword(data["password"])
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	isAmbasador := strings.Contains(c.Path(), "/api/ambassador")

	var scope string

	if isAmbasador {
		scope = "ambassador"
	} else {
		scope = "admin"
	}

	// prevent ambassador users to login into admin dashboard
	// admins can login into ambassador dashboard
	if !isAmbasador && user.IsAmbasador {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthorized.",
		})
	}

	token, err := middlewares.GenerateJWT(user.Id, scope)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

// GetUser gets credentials from signed in user by ID
func GetUser(c *fiber.Ctx) error {
	id, _ := middlewares.GetUserId(c)

	var user models.User

	database.DB.Where("id = ?", id).First(&user)

	if strings.Contains(c.Path(), "/api/ambassador") {
		ambassador := models.Ambassador(user)
		ambassador.CalculateRevenue(database.DB)
		return c.JSON(ambassador)
	}

	return c.JSON(user)
}

// Logout signout user and removes cookie
func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

// UpdateUserInfo updates user info in db
func UpdateUserInfo(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	id, _ := middlewares.GetUserId(c)

	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
	}

	user.Id = id

	database.DB.Model(models.User{}).Where("id = ?", id).Updates(&user)

	return c.JSON(user)
}

// UpdatePassword updates user password
func UpdatePassword(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}

	id, _ := middlewares.GetUserId(c)

	user := models.User{}

	user.Id = id

	user.SetPassword(data["password"])

	database.DB.Model(models.User{}).Where("id = ?", id).Updates(&user)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
