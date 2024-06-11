package main

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type Car struct {
	Name  string
	Price float64
}

var cars []Car

func generateCars() {
	cars = append(cars, Car{Name: "Mercedes", Price: 8700})
	cars = append(cars, Car{Name: "Fiat", Price: 6500})
	cars = append(cars, Car{Name: "Corsa", Price: 3500})
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
	saveCar(*car)
	return C.JSON(200, cars)
}

func saveCar(car Car) error {
	db, err := sql.Open("sqlite3", "cars.db")
	if err != nil {
		return err
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO cars (Name, Price) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(car.Name, car.Price)
	if err != nil {
		return err
	}
	return nil
}
