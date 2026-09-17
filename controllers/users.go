package controllers

import (
	"strconv"

	"github.com/detectivegot/fiber-learn/data"
	"github.com/detectivegot/fiber-learn/models"
	"github.com/gofiber/fiber/v3"
)

func GetUser(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid user id")
	}
	for i := range data.Users {
		if data.Users[i].Id==id {
			return c.JSON(data.Users[i])
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("User " + idStr + " is not found.")
}

func GetUsers(c fiber.Ctx) error {
	return c.JSON(data.Users)
}

func CreateUser(c fiber.Ctx) error {
	var user models.Users

	if err := c.Bind().Body(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invaild request.")
	}

	data.Users = append(data.Users, user)
	return c.Status(fiber.StatusCreated).JSON(user)
}

func UpdateUser(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid user id")
	}
	var user models.Users
	if err := c.Bind().Body(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invaild request.")
	}
	for i := range data.Users {
		if id == data.Users[i].Id {
			data.Users[i] = user
			return c.Status(fiber.StatusAccepted).JSON(user)
		}
	}
	return c.Status(fiber.StatusBadRequest).SendString("Invalid user")
}

func DeleteUser(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid user id")
	}
	for i := range data.Users {
		if id == data.Users[i].Id {
			user := data.Users[i]
			data.Users = append(data.Users[:i], data.Users[i+1:]...)
			return c.Status(fiber.StatusAccepted).JSON(user)
		}
	}
	return c.Status(fiber.StatusBadRequest).SendString("Invalid user")
}