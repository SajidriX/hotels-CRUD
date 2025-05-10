package main

import (
	"hotels/hotels"
	"hotels/users"
	"net/http"

	"github.com/labstack/echo/v4"
)

func hi(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"message": "welcome!"})
}

func main() {

	if err := users.InitDB(); err != nil {
		panic("Failed to init users DB")
	}
	if err := hotels.InitDB(); err != nil {
		panic("Failed to init hotels DB")
	}

	e := echo.New()

	e.POST("/auth", users.CreateUser)

	e.POST("/hotelsCreate", hotels.CreateHotel)
	e.GET("/hotels", hotels.GetHotels)
	e.PATCH("/hotelsPatch/:name", hotels.PatchHotels)
	e.DELETE("/hotelsDelete/:name", hotels.DeleteHotels)
	e.GET("/hotelsByCoun/:country", hotels.GetHotelByCountry)
	e.GET("hotelsByName/:name", hotels.GetHotelByName)
	e.GET("/", hi)

	e.Logger.Fatal(e.Start(":1228"))
}
