//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

type Selector interface {
	SelectLiftBySize()
}

type Vehicle struct {
	size      string
	modelName string
}

func (v Vehicle) SelectLiftBySize() {
	if v.size == "motorcicle" {
		fmt.Printf("%v is the vehicle type and %v is her model name, use small lifts \n", v.size, v.modelName)
	} else if v.size == "car" {
		fmt.Printf("%v is the vehicle type and %v is her model name, use standad lifts \n", v.size, v.modelName)
	} else {
		fmt.Printf("%v is the vehicle and %v is her model name, use large lifts \n", v.size, v.modelName)
	}

}

func process(vehicles []Selector) {
	fmt.Println("Proccesing:")
	for i := 0; i < len(vehicles); i++ {
		vehicle := vehicles[i]
		fmt.Printf("--Vehicle: %v--\n", vehicle)
		vehicle.SelectLiftBySize()
	}
	fmt.Println()
}

func main() {
	vehicles := []Selector{
		Vehicle{size: "motorcicle", modelName: "a400m"},
		Vehicle{size: "truck", modelName: "5000"},
		Vehicle{size: "car", modelName: "A4a1"},
	}
	process(vehicles)
}
