package routes

import (
	"github.com/detectivegot/fiber-learn/controllers"
	"github.com/gofiber/fiber/v3"
)

func UserRoutes(v1 fiber.Router) {
	v1.Route("/users", func (r fiber.Router) {
		r.Get("/", controllers.GetUsers)
		r.Get("/:id<int>", controllers.GetUser)
		r.Post("/", controllers.CreateUser)
		r.Put("/:id", controllers.UpdateUser)
		r.Delete("/:id", controllers.DeleteUser)
	})
}