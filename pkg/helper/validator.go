package helper

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	Validator *validator.Validate
}

// Validate is a function that validates the given data using the CustomValidator.
//
// It takes a data interface{} as a parameter and returns an error as validationErrors.
func (v *CustomValidator) Validate(data any) (validationErrors error) {

	errs := v.Validator.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			validationErrors = errors.Join(validationErrors, err)
		}
	}
	validationErrors = echo.NewHTTPError(http.StatusBadRequest, validationErrors)
	return
}
