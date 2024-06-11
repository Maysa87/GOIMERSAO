package main

import (
	"github.com/labstack/echo/v4"
)

type Car struct {
	name  string
	price float64
}

var cars []Car

func generateCars() {
	cars = append(cars, Car{name: "Mercedes", price: 870000})
	cars = append(cars, Car{name: "Fiat", price: 65000})
	cars = append(cars, Car{name: "Corsa", price: 35000})
}

func main() {
	generateCars()
	e := echo.New()
	e.GET("/cars", getCars)
	e.POST("/cars", createCar)
	e.Logger.Fatal(e.Start(":8080"))
}

func getCars(C echo.Context) error {
	return C.JSON(200, cars)
}

func createCar(C echo.Context) error {
	car := new(Car)
	if err := C.Bind(car); err != nil {
		return err
	}
	cars = append(cars, *car)
	return C.JSON(200, cars)
}
