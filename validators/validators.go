package validators

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateOkayTitle(field validator.FieldLevel) bool {
	return strings.Contains(field.Field().String(), "okay")
}