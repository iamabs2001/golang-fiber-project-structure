package main

import (
	"fmt"
	"phonebook/src/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("./src/views/public", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	// Index / Master / Main Controller
	controllers.SetupRoutes(app)
	fmt.Println("Server running at port : 8080")
	app.Listen(":8080")
}
