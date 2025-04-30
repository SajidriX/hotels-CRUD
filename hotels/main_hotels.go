package hotels

import (
	"gorm.io/driver/sqlite"
	// "github.com/labstack/echo/v4"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Hotel struct {
	gorm.Model
	Name        string `json:"name" validate:"required,min=3,max=30"`
	Country     string `json:"country" validate:"required,min=3,max=30"`
	Description string `json:"description" validate:"required,min=3,max=250"`
}

type hotelGet struct {
	Name        string `json:"name" validate:"required,min=3,max=30"`
	Country     string `json:"country" validate:"required,min=3,max=30"`
	Description string `json:"description" validate:"required,min=3,max=250"`
}

var valid = validator.New()

func toHotelGet(h *Hotel) hotelGet {
	return hotelGet{
		Name:        h.Name,
		Country:     h.Country,
		Description: h.Description,
	}
}

var db *gorm.DB

func CreateHotel(c echo.Context) error {
	hotel := new(Hotel)
	if err := c.Bind(hotel); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Bad input"})
	}

	if err := valid.Struct(hotel); err != nil {
		return c.JSON(422, echo.Map{"validation errror": err.Error()})
	}

	if err := db.Create(hotel).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error making hotel"})
	}

	return c.JSON(http.StatusOK, hotel)
}

func GetHotels(c echo.Context) error {
	var hotels []Hotel

	if err := db.Find(&hotels).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error getting hotel, we are sorry"})
	}

	var response []hotelGet

	for _, ht := range hotels {
		response = append(response, toHotelGet(&ht))
	}

	return c.JSON(http.StatusOK, response)

}

func PatchHotels(c echo.Context) error {
	name := c.Param("name")
	var hotel Hotel

	if err := db.Where("name = ?", name).First(&hotel).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "hotel not found"})
	}

	var updateData map[string]interface{}

	if err := c.Bind(&updateData); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}
	if err := db.Model(&hotel).Updates(updateData).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to update cheese"})
	}

	return c.JSON(http.StatusOK, hotel)
}

func DeleteHotels(c echo.Context) error {
	name := c.Param("name")
	var hotel Hotel

	if err := db.Where("name = ?", name).First(&hotel).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "hotel not found"})
	}

	if err := db.Delete(&hotel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to delete hotel, sorry"})
	}
	return c.JSON(http.StatusOK, echo.Map{"Hotel deleted": hotel})
}

func InitDB() error {
	var err error
	db, err = gorm.Open(sqlite.Open("hotels.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	return db.AutoMigrate(&Hotel{})
}

func GetHotelByCountry(c echo.Context) error {
	country := c.Param("country")
	var hotels []Hotel

	if err := db.Where("country = ?", country).Find(&hotels).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to filter hotels, sorry"})
	}

	var response []hotelGet

	for _, ht := range hotels {
		response = append(response, toHotelGet(&ht))
	}

	return c.JSON(http.StatusOK, response)
}

func GetHotelByName(c echo.Context) error {
	name := c.Param("name")
	var hotels []Hotel

	if err := db.Where("name = ?", name).Find(&hotels).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to filter hotels by name, sorry"})
	}

	var resp []hotelGet
	for _, hotel := range hotels {
		resp = append(resp, toHotelGet(&hotel))
	}

	return c.JSON(http.StatusOK, resp)
}
