package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func GetContact(c *fiber.Ctx) error {
	return c.SendString("Hello, " + c.Params("mobileNo"))
}
