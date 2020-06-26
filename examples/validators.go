package example

import (
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type (
	// CustomValidator for the example
	CustomValidator struct {
		validator *validator.Validate
	}

	// ExampleModelValidator for example
	ExampleModelValidator struct {
		gorm.Model
		ExampleID    int64  `validate:"required"`
		Pin      int64  `validate:"required"`
		Customer string `validate:"required"`
		Quantity int64  `validate:"required"`
		Notes    string
	}
)

// Validate the example route
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// RegisterValidator for the example
func RegisterValidator(e *echo.Echo) {
	e.Validator = &CustomValidator{validator: validator.New()}
}
