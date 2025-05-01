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

	e.POST("/hotels", hotels.CreateHotel)
	e.GET("/hotels", hotels.GetHotels)
	e.PATCH("/hotels/:name", hotels.PatchHotels)
	e.DELETE("/hotels/:name", hotels.DeleteHotels)
	e.GET("/hotels/:country", hotels.GetHotelByCountry)
	e.GET("hotels/:name", hotels.GetHotelByName)

	e.Logger.Fatal(e.Start(":1228"))
}
