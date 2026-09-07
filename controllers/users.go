package controllers

import (
	"github.com/gofiber/fiber/v3"
)

func GetUser(c fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("Get one user " + id)
}

func GetUsers(c fiber.Ctx) error {
	return c.SendString("Get users")
}

func CreateUser(c fiber.Ctx) error {
	return c.SendString("Create user")
}

func UpdateUser(c fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("Update user " + id)
}

func DeleteUser(c fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("Delete user " + id)
}