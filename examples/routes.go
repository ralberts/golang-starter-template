package example

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo/v4"
)

var DB *gorm.DB

// Register the main input to for a example route
func Register(group string, router *echo.Echo) error {
	route := group + "/example"
	log.Print("Register route: " + route)
	RegisterValidator(router)
	router.GET(route, Registration)
	router.POST(route, Update)
	// router.POST("/login", UsersLogin)
	return nil
}

// Registration example registration for the default route
func Registration(c echo.Context) (err error) {
	u := new(ExampleModel)
	if err = c.Bind(u); err != nil {
		return
	}

	return c.JSON(http.StatusOK, u)
}

// Update for adding a job
func Update(c echo.Context) (err error) {
	j := new(ExampleModel)
	if err = c.Bind(j); err != nil {
		return
	}

	v := ExampleModelValidator(*j)
	if err = c.Validate(v); err != nil {
		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println("Namespace:", err.Namespace())
			fmt.Println("Field:", err.Field())
			fmt.Println("StructNamespace:", err.StructNamespace())
			fmt.Println("StructField:", err.StructField())
			fmt.Println("Tag:", err.Tag())
			fmt.Println("ActualTag:", err.ActualTag())
			fmt.Println("Kind:", err.Kind())
			fmt.Println("Type:", err.Type())
			fmt.Println("Value:", err.Value())
			fmt.Println("Param:", err.Param())
			fmt.Println()
		}
	}

	DB.NewRecord(j)
	DB.Create(j)

	return c.JSON(http.StatusOK, j)
}
