package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nezbut/hello-http-go/handlers"
)

func main() {
	app := fiber.New()
	app.Get("/", handlers.MainRoute)
	log.Fatal(app.Listen(":8000"))
}
