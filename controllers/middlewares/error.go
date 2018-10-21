package middlewares

import (
	"net/http"
	"strings"

	"github.com/PrinceNorin/monga/utils/errors"
	"github.com/PrinceNorin/monga/utils/messages"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, ginerr := range c.Errors {
			err := ginerr.Err
			if strings.Contains(err.Error(), "violates unique constraint") {
				err = errors.ErrNotUnique
			}
			status := getStatus(err)
			payload := gin.H{"status": status}

			if msg, ok := err.(*messages.Messages); ok {
				payload["messages"] = msg.GetAllErrors()
			}
			if status == http.StatusInternalServerError ||
				status == http.StatusBadRequest {
				// TODO: send to error logging service, sentry maybe?
				// captureError(err)
			} else {
				if err == gorm.ErrRecordNotFound {
					err = errors.ErrRecordNotFound
				}
				payload["code"] = err.Error()
			}
			c.JSON(status, payload)
		}
	}
}

func getStatus(err error) int {
	switch err {
	case errors.ErrRecordNotFound,
		gorm.ErrRecordNotFound:
		return http.StatusNotFound
	case errors.ErrBadRequest:
		return http.StatusBadRequest
	case errors.ErrValidation, errors.ErrNotUnique:
		return http.StatusUnprocessableEntity
	case errors.ErrUnauthorized:
		return http.StatusUnauthorized
	default:
		if _, ok := err.(*messages.Messages); ok {
			return http.StatusUnprocessableEntity
		}
		return http.StatusInternalServerError
	}
}
