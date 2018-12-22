package validators

import (
	"fmt"
	"strings"

	"github.com/PrinceNorin/monga/utils/messages"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/iancoleman/strcase"
	validator "gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("date", DateValidatorFunc)
}

func Bind(v interface{}, c *gin.Context) error {
	contentType := c.ContentType()
	if strings.Contains(contentType, "multipart/form-data") {
		return c.ShouldBindWith(v, binding.FormMultipart)
	} else {
		return c.ShouldBindWith(v, binding.JSON)
	}
}

func Validate(f interface{}, msg *messages.Messages) error {
	err := validate.Struct(f)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fieldName := strcase.ToLowerCamel(err.Field())
			if err.Param() != "" {
				msg.AddError(fieldName, fmt.Errorf("errors_%s: %s", err.Tag(), err.Param()))
			} else {
				msg.AddError(fieldName, fmt.Errorf("errors_%s", err.Tag()))
			}
		}
		return err
	}
	return nil
}
