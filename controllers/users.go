package controllers

import (
	"database/sql"

	"github.com/detectivegot/fiber-learn/models"
	"github.com/gofiber/fiber/v3"
)

// func GetUser(c fiber.Ctx) error {
// 	idStr := c.Params("id")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).SendString("Invalid user id")
// 	}
// 	for i := range data.Users {
// 		if data.Users[i].Id==id {
// 			return c.JSON(data.Users[i])
// 		}
// 	}
// 	return c.Status(fiber.StatusNotFound).SendString("User " + idStr + " is not found.")
// }

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

// func CreateUser(c fiber.Ctx) error {
// 	var user models.Users

// 	if err := c.Bind().Body(&user); err != nil {
// 		var validateErrs validator.ValidationErrors
// 		if errors.As(err, &validateErrs) {
// 			out := make([]fiber.Map, 0, len(validateErrs))
// 			for _, e := range validateErrs {
// 				out = append(out, fiber.Map{
// 					"field": e.Field(),
// 					"rule": e.Tag(),
// 					"param": e.Param(),
// 					"value": e.Value(),
// 				})
// 			}
// 			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": out})
// 		}
// 		return c.Status(fiber.StatusBadRequest).SendString("Invalid user")
// 	}

// 	data.Users = append(data.Users, user)
// 	return c.Status(fiber.StatusCreated).JSON(user)
// }

// func UpdateUser(c fiber.Ctx) error {
// 	idStr := c.Params("id")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).SendString("Invalid user id")
// 	}
// 	var user models.Users
// 	if err := c.Bind().Body(&user); err != nil {
// 		var validateErrs validator.ValidationErrors
// 		if errors.As(err, &validateErrs) {
// 			out := make([]fiber.Map, 0, len(validateErrs))
// 			for _, e := range validateErrs {
// 				out = append(out, fiber.Map{
// 					"field": e.Field(),
// 					"rule": e.Tag(),
// 					"param": e.Param(),
// 					"value": e.Value(),
// 				})
// 			}
// 			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": out})
// 		}
// 		return c.Status(fiber.StatusBadRequest).SendString("Invalid user")
// 	}
// 	for i := range data.Users {
// 		if id == data.Users[i].Id {
// 			data.Users[i] = user
// 			return c.Status(fiber.StatusAccepted).JSON(user)
// 		}
// 	}
// 	return c.Status(fiber.StatusBadRequest).SendString("Invalid user")
// }

// func DeleteUser(c fiber.Ctx) error {
// 	idStr := c.Params("id")
// 	id, err := strconv.Atoi(idStr)

// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).SendString("Invalid user id")
// 	}
// 	for i := range data.Users {
// 		if id == data.Users[i].Id {
// 			user := data.Users[i]
// 			data.Users = append(data.Users[:i], data.Users[i+1:]...)
// 			return c.Status(fiber.StatusAccepted).JSON(user)
// 		}
// 	}
// 	return c.Status(fiber.StatusBadRequest).SendString("Invalid user")
// }