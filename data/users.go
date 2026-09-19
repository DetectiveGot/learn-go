package data

import (
	"uuid"

	"github.com/detectivegot/fiber-learn/models"
)

var Users = []models.Users{
	{
		Id: uuid.New(),
		Name: "Got",
		Age: 10,
	},
	{
		Id: uuid.New(),
		Name: "Laplace",
		Age: 11,
	},
	{
		Id: uuid.New(),
		Name: "Orivia",
		Age: 13,
	},
}