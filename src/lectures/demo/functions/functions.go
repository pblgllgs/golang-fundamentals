package main

import "fmt"

func double(x int) int {
	return x * x
}

func add(lhs, rhs int) int {
	return lhs + rhs
}

func greet() {
	fmt.Println("Hello from greet function")
}

func main() {

	greet()

	doblar := double(2)
	fmt.Println(doblar)

	agregar := add(doblar, 2)
	fmt.Println(agregar)

	anotherFunction := add(double(3), 1)
	fmt.Println("El resultado es", anotherFunction)

}
