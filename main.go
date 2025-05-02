package main

import (
	"hotels/hotels"
	"hotels/users"

	"github.com/labstack/echo/v4"
)

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

	e.Logger.Fatal(e.Start(":1228"))
}
