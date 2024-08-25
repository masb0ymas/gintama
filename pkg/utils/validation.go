package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

var validate = validator.New()

/*
Private Function
*/
func validateStruct(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

func Validate(data interface{}) (code int32, message string, errors []string) {
	if errs := validateStruct(data); len(errs) > 0 && errs[0].Error {
		errMsgs := make([]string, 0)

		for _, err := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				err.FailedField,
				err.Value,
				err.Tag,
			))
		}

		return http.StatusBadRequest, "Bad Request", errMsgs
	}

	return 200, "", nil
}

// parse form data body
func ParseFormData(c *gin.Context, body interface{}) (code int32, message string, errors []string) {
	if err := c.ShouldBindJSON(body); err != nil {
		errMsgs := make([]string, 0)
		errMsgs = append(errMsgs, "Unprocessable Entity")

		fmt.Println(body, err)

		return http.StatusUnprocessableEntity, "Unprocessable Entity", errMsgs
	}

	return 200, "", nil
}

// parse form data body and validate form
func ParseFormDataAndValidate(c *gin.Context, body interface{}) (code int32, message string, errors []string) {
	if code, message, errors := ParseFormData(c, body); errors != nil {
		return code, message, errors
	}

	return Validate(body)
}
