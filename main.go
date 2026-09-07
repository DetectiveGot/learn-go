package main

import (
	"fmt"

	"github.com/detectivegot/fiber-learn/routes"
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()
	
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Welcome to Fiber")
	})
	routes.HandleRoutes(app)
	app.Listen(":8089")
	fmt.Println("Server is running at port 8089")
}