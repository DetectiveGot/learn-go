package main

import (
	"fmt"

	"github.com/detectivegot/fiber-learn/routes"
	"github.com/detectivegot/fiber-learn/validators"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/html/v3"
)



func main() {

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
		StructValidator: validators.New(),
	})
	
	app.Get("/", func(c fiber.Ctx) error {
		render := c.Render("index", fiber.Map{
			"Title": "Eman",
			"Text": "this is learning go-fiber!",
		})
		return render
	})

	routes.HandleRoutes(app)
	app.Listen(":8089")
	fmt.Println("Server is running at port 8089")
}