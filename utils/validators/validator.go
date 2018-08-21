package validators

import (
	"fmt"

	"github.com/PrinceNorin/monga/utils/messages"
	"github.com/iancoleman/strcase"
	validator "gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func Validate(f interface{}, msg *messages.Messages) error {
	err := validate.Struct(f)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fieldName := strcase.ToLowerCamel(err.Field())
			if err.Tag() == "unique" {
				msg.AddError(fieldName, "errors_unique")
			} else if err.Param() != "" {
				msg.AddError(fieldName, fmt.Sprintf("errors_%s: %s", err.Tag(), err.Param()))
			} else {
				msg.AddError(fieldName, fmt.Sprintf("errors_%s", err.Tag()))
			}
		}
		return err
	}
	return nil
}
