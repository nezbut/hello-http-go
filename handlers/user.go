package handlers

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
)

func MainRoute(c *fiber.Ctx) error {
	var col = color.New(color.BgCyan).Add(color.BgRed)
	str := fmt.Sprintf("Hello Nezbut from HTTP World, ip: %s", c.IP())
	col.Println(str)
	return c.SendString(str)
}
