package routes

import (
	"ecommerce/src/controllers"
	"ecommerce/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")

	admin := api.Group("admin")

	admin.Post("register", controllers.Register)
	admin.Post("login", controllers.Login)

	adminAuthenticated := admin.Use(middlewares.IsAuthenticated)
	adminAuthenticated.Get("user", controllers.GetUser)
	adminAuthenticated.Post("logout", controllers.Logout)
	adminAuthenticated.Put("user/info", controllers.UpdateUserInfo)
	adminAuthenticated.Put("user/password", controllers.UpdatePassword)
	adminAuthenticated.Get("ambasadors", controllers.GetAmbasadors)
	adminAuthenticated.Get("products", controllers.GetProducts)
	adminAuthenticated.Post("products", controllers.CreateProduct)
	adminAuthenticated.Get("products/:id", controllers.GetProduct)
	adminAuthenticated.Put("products/:id", controllers.UpdateProduct)
	adminAuthenticated.Delete("products/:id", controllers.DeleteProduct)
	adminAuthenticated.Get("users/:id/links", controllers.GetLinksByUserId)
	adminAuthenticated.Get("orders", controllers.GetAllOrders)
}
