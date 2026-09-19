package routes

import (
	"database/sql"

	"github.com/detectivegot/fiber-learn/controllers"
	"github.com/gofiber/fiber/v3"
)

func UserRoutes(v1 fiber.Router, db *sql.DB) {
	v1.Route("/users", func (r fiber.Router) {
		r.Get("/", controllers.GetUsers(db))
		r.Get("/:id", controllers.GetUser(db))
		r.Post("/", controllers.CreateUser(db))
		r.Put("/:id", controllers.UpdateUser(db))
		r.Delete("/:id", controllers.DeleteUser(db))
	})
}