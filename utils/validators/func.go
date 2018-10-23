package validators

import (
	"time"

	validator "gopkg.in/go-playground/validator.v9"
)

func DateValidatorFunc(fl validator.FieldLevel) bool {
	v := fl.Field().Interface()
	if v == nil {
		return false
	}
	d, ok := v.(time.Time)
	if !ok {
		return false
	}
	if d.Format("2006-01-02") == "0001-01-01" {
		return false
	}
	return true
}
