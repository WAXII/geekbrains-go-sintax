package main

import (
	"fmt"
	"vehicles"
)

func main() {
	fmt.Println(vehicles.GetAllVehicles())
	vehicles.AddCar("v1", "Mazda", 2015, 162.5)
	vehicles.AddCar("v2", "Scania", 2014, 3852.7)
	fmt.Println(vehicles.GetAllVehicles())
	vehicles.OpenCloseWindow("v1", true)
	fmt.Println(vehicles.GetAllVehicles())
	vehicles.StartStopEngine("v2", true)
	fmt.Println(vehicles.GetAllVehicles())
}
