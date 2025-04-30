package users

import (
	"net/http"

	//"hotels/hotels"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"name" validate:"required,min=3,max=30"`
}

var db *gorm.DB
var validate = validator.New()

func CreateUser(c echo.Context) error {
	user := new(User)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Bad input"})
	}

	if err := validate.Struct(user); err != nil {
		return c.JSON(422, echo.Map{"error": "validation error"})
	}

	if err := db.Create(user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error creating user"})
	}

	return c.JSON(http.StatusOK, user)
}

func InitDB() error {
	var err error
	db, err = gorm.Open(sqlite.Open("hotels.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	return db.AutoMigrate(&User{})
}
