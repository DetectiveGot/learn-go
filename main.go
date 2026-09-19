package main

import (
	"fmt"
	"log"

	"github.com/detectivegot/fiber-learn/config"
	"github.com/detectivegot/fiber-learn/database"
	"github.com/detectivegot/fiber-learn/routes"
	"github.com/detectivegot/fiber-learn/validators"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/html/v3"
	"github.com/joho/godotenv"
)


func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Couldn't load .envL ", err)
	}

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
		StructValidator: validators.New(),
	})

	cfg := config.Load()
	db, err := database.Connect(cfg.DatabaseURL)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("DB is connected successfully.")
	
	app.Get("/", func(c fiber.Ctx) error {
		render := c.Render("index", fiber.Map{
			"Title": "Eman",
			"Text": "this is learning go-fiber!",
		})
		return render
	})

	routes.HandleRoutes(app, db)
	app.Listen(":8089")
	fmt.Println("Server is running at port 8089")
}