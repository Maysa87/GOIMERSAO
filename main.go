package main

import (
	"github.com/labstack/echo/v4"
)

type Car struct {
	name  string
	price float64
}

var cars []Car

func createCars() {
	cars = append(cars, Car{name: "Mercedes", price: 870000})
	cars = append(cars, Car{name: "Fiat", price: 65000})
	cars = append(cars, Car{name: "Corsa", price: 35000})
}

func getCars(C echo.Context) error {
	return C.JSON(200, cars)
}
func main() {
	createCars()
	e := echo.New()
	e.GET("/cars", getCars)
	e.Logger.Fatal(e.Start(":8080"))
}
