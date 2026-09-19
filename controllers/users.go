package controllers

import (
	"database/sql"
	"uuid"

	"errors"

	"github.com/detectivegot/fiber-learn/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GetUser(db *sql.DB) fiber.Handler {
	return func (c fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid user id")
		}
		var user models.Users
		err = db.QueryRow("SELECT id, name, age FROM users WHERE id = $1", id).Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		
		return c.Status(fiber.StatusOK).JSON(user)
	}
}

func GetUsers(db *sql.DB) fiber.Handler {
	return func (c fiber.Ctx) error {
		rows, err := db.Query("SELECT id, name, age FROM users LIMIT 5")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		defer rows.Close()
		result := make([]models.Users, 0)
		for rows.Next() {
			var user models.Users
			err := rows.Scan(&user.Id, &user.Name, &user.Age)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			}
			result = append(result, user)
		}

		if err := rows.Err(); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.Status(fiber.StatusOK).JSON(result)
	}
}

func CreateUser(db *sql.DB) fiber.Handler {
	return func (c fiber.Ctx) error {
		var user models.Users

		if err := c.Bind().Body(&user); err != nil {
			var validateErrs validator.ValidationErrors
			if errors.As(err, &validateErrs) {
				out := make([]fiber.Map, 0, len(validateErrs))
				for _, e := range validateErrs {
					out = append(out, fiber.Map{
						"field": e.Field(),
						"rule": e.Tag(),
						"param": e.Param(),
						"value": e.Value(),
					})
				}
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": out})
			}
			return c.Status(fiber.StatusBadRequest).SendString("Invalid user")
		}

		err := db.QueryRow("INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id, name, age",
			user.Name,
			user.Age,
		).Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.Status(fiber.StatusCreated).JSON(user)
	}
}

func UpdateUser(db *sql.DB) fiber.Handler {
	return func (c fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid user id")
		}
		var user models.Users
		if err := c.Bind().Body(&user); err != nil {
			var validateErrs validator.ValidationErrors
			if errors.As(err, &validateErrs) {
				out := make([]fiber.Map, 0, len(validateErrs))
				for _, e := range validateErrs {
					out = append(out, fiber.Map{
						"field": e.Field(),
						"rule": e.Tag(),
						"param": e.Param(),
						"value": e.Value(),
					})
				}
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": out})
			}
			return c.Status(fiber.StatusBadRequest).SendString("Invalid user")
		}

		err = db.QueryRow("UPDATE users SET name=$1, age=$2 WHERE id=$3 RETURNING id, name, age",
			user.Name,
			user.Age,
			id,
		).Scan(&user.Id, &user.Name, &user.Age)
		
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		
		return c.Status(fiber.StatusOK).JSON(user)
	}
}

func DeleteUser(db *sql.DB) fiber.Handler {
	return func (c fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := uuid.Parse(idStr)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid user id")
		}
		var user models.Users
		err = db.QueryRow("DELETE FROM users WHERE id=$1 RETURNING id, name, age", id).Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.Status(fiber.StatusOK).JSON(user)
	}
}