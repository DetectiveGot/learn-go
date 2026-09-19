package models

import "uuid"

type Users struct {
	Id   uuid.UUID   `json:"id"`
	Name string `json:"name" validate:"required"`
	Age  int16    `json:"age" validate:"required,gte=0,lte=100"`
}