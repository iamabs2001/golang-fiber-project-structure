package controllers

import (
	"github.com/gofiber/fiber/v2"
)

// Main index route
func SetupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)
	app.Get("/view", viewHtml)
	// Contact Routes
	app.Get("/getContacts", helloWorld)
	app.Get("/getContact/:mobileNo", GetContact)
	app.Post("/saveContact", helloWorld)
	app.Put("/updateContact", helloWorld)
	app.Delete("/deleteContact", helloWorld)
}

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, Phonebook!")
}

func viewHtml(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}
