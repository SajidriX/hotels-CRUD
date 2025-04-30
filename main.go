package main

import (
	"hotels/hotels"
	"hotels/users"

	"github.com/labstack/echo/v4"
)

func main() {
	// Инициализация БД
	if err := users.InitDB(); err != nil {
		panic("Failed to init users DB")
	}
	if err := hotels.InitDB(); err != nil {
		panic("Failed to init hotels DB")
	}

	e := echo.New()

	// Роуты пользователей
	e.POST("/auth", users.CreateUser)

	// Роуты отелей
	e.POST("/hotels", hotels.CreateHotel) // или /hotels_make, если так задумано
	e.GET("/hotels", hotels.GetHotels)
	e.PATCH("/hotels/:name", hotels.PatchHotels)
	e.DELETE("/hotels/:name", hotels.DeleteHotels)
	e.GET("/hotels/:country", hotels.GetHotelByCountry)
	e.GET("hotels/:name", hotels.GetHotelByName)

	e.Logger.Fatal(e.Start(":1228"))
}
