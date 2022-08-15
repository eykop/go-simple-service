package data

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func validateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile("[a-z]+-[a-z]+-[a-z]+")
	return len(re.FindAllString(fl.Field().String(), -1)) == 1
}
