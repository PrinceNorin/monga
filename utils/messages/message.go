package messages

import (
	"github.com/PrinceNorin/monga/utils/errors"
	"github.com/gin-gonic/gin"
)

type Messages struct {
	c      *gin.Context
	errors map[string][]string
}

const messageKey = "monga.messages"

func GetMessages(c *gin.Context) *Messages {
	if m, ok := c.Get(messageKey); ok {
		msg := m.(*Messages)
		msg.c = c
		return msg
	}
	msg := &Messages{c, make(map[string][]string)}
	c.Set(messageKey, msg)
	return msg
}

func (msg *Messages) AddError(key string, err error) {
	msg.errors[key] = append(msg.errors[key], err.Error())
	msg.setInContext()
}

func (msg *Messages) Error() string {
	return errors.ErrValidation.Error()
}

func (msg *Messages) GetAllErrors() map[string][]string {
	return msg.errors
}

func (msg *Messages) GetError(key string) []string {
	if errs, ok := msg.errors[key]; ok {
		return errs
	}
	return nil
}

func (msg *Messages) HasErrors() bool {
	return len(msg.errors) > 0
}

func (msg *Messages) HasError(key string) bool {
	return len(msg.GetError(key)) > 0
}

func (msg *Messages) Clear() {
	msg.errors = make(map[string][]string)
}

func (msg *Messages) ClearError(key string) {
	msg.errors[key] = nil
}

func (msg *Messages) setInContext() {
	msg.c.Set(messageKey, msg)
}
