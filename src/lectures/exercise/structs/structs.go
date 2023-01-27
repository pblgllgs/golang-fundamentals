//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Coordinate struct {
	x, y int
}

type Rectangulo struct {
	a Coordinate
	b Coordinate
}

func with(rect Rectangulo) int {
	return rect.b.x - rect.a.x
}

func length(rect Rectangulo) int {
	return rect.a.y - rect.b.y
}

func area(rect Rectangulo) int {
	return with(rect) * length(rect)
}

func perameter(rect Rectangulo) int {
	return with(rect)*2 + length(rect)*2
}

func printInfo(rect Rectangulo) {
	fmt.Println("The area is", area(rect))
	fmt.Println("The perimeter is", perameter(rect))
}

func main() {
	rectangulo := Rectangulo{a: Coordinate{0, 7}, b: Coordinate{10, 0}}
	printInfo(rectangulo)
	rectangulo = Rectangulo{a: Coordinate{0, 14}, b: Coordinate{20, 0}}
	printInfo(rectangulo)
}
