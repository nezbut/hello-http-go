package handlers

import "github.com/gofiber/fiber/v3"

func MainRoute(c fiber.Ctx) error {
	return c.SendString("Hello Nezbut from HTTP World")
}
