package apperror

import (
	"net/http"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

// APPError is the default error struct containing detailed information about the error
type APPError struct {
	// HTTP Status code to be set in response
	Status int `json:"-"`
	// Message is the error message that may be displayed to end users
	Message string `json:"message,omitempty"`
	// Meta is the error detail detail data
	Meta interface{} `json:"meta,omitempty"`
}

// NewStatus generates new error containing only http status code
func NewStatus(status int) *APPError {
	return &APPError{Status: status}
}

// New generates an application error
func New(status int, msg string, meta ...interface{}) *APPError {
	err := &APPError{Status: status, Message: msg}
	if len(meta) > 0 {
		err.Meta = meta
	}
	return err
}

// Error returns the error message.
func (e APPError) Error() string {
	return e.Message
}

// Error Struct Validate
type Error struct {
	Errors map[string]interface{} `json:"message"`
}

// Response writes an error response to client
func Response(c *gin.Context, err error) {
	switch err.(type) {
	case *APPError:
		e := err.(*APPError)
		if e.Message == "" {
			c.AbortWithStatus(e.Status)
		} else {
			c.AbortWithStatusJSON(e.Status, e)
		}
		return
	case validation.Errors:
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, err)
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
}
