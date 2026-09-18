package models

type Users struct {
	Id   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
	Age  int    `json:"age" validate:"required,gte=0,lte=100"`
}