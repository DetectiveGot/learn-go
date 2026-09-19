package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v3"
)

func HandleRoutes(app *fiber.App, db *sql.DB) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	UserRoutes(v1, db);
}