package validators

import (
	"github.com/go-playground/validator/v10"
)
type StructValidator struct {
	validate *validator.Validate
}

func (v *StructValidator) Validate(out any) error {
	return v.validate.Struct(out)
}

func New() *StructValidator {
	return &StructValidator{
		validate: validator.New(),
	}
}