package main

import (
	"log"

	"github.com/eyupduran/fiber-api/database"
	"github.com/eyupduran/fiber-api/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome my go api")
}

func setupRoutes(app *fiber.App) {
	//welcome endpoint
	app.Get("/api", welcome)
	//user endpoint
	app.Post("/api/addUser", routes.CreateUser)
	app.Get("/api/getAllUser", routes.GetUsers)
	app.Get("/api/getUserById/:id", routes.GetUser)
	app.Put("/api/updateUserById/:id", routes.UpdateUser)
	app.Delete("/api/deleteUserById/:id", routes.DeleteUser)
	//product endpoint
	app.Post("/api/addProduct", routes.CreateProduct)
	app.Get("/api/getAllProduct", routes.GetProducts)
	app.Get("/api/getProductById/:id", routes.GetProduct)
	app.Put("/api/updateProductById/:id", routes.UpdateProduct)

	// Order endpoints
	app.Post("/api/addOrder", routes.CreateOrder)
	app.Get("/api/getAllOrder", routes.GetOrders)
	app.Get("/api/getOrderById/:id", routes.GetOrder)

}
func main() {
	database.ConnectDb()
	app := fiber.New()
	setupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
