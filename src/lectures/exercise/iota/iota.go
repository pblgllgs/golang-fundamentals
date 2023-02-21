//--Summary:
//  Create a calculator that can perform basic mathematical operations.
//
//--Requirements:
//* Mathematical operations must be defined as constants using iota
//* Write a receiver function that performs the mathematical operation
//  on two operands
//* Operations required:
//  - Add, Subtract, Multiply, Divide
//* The existing function calls in main() represent the API and cannot be changed
//
//--Notes:
//* Your program is complete when it compiles and prints the correct results

package main

import "fmt"

type Operation int

const (
	Add Operation = iota
	Sub
	Mul
	Div
)

func (c Operation) calculate(a, b int) int {
	switch c {
	case Add:
		return a + b
	case Sub:
		return a - b
	case Mul:
		return a * b
	case Div:
		return a / b
	}
	panic("unhandled operation")
}

func main() {
	add := Operation(Add)
	fmt.Println(add.calculate(2, 2)) // = 4

	sub := Operation(Sub)
	fmt.Println(sub.calculate(10, 3)) // = 7

	mul := Operation(Mul)
	fmt.Println(mul.calculate(3, 3)) // = 9

	div := Operation(Div)
	fmt.Println(div.calculate(100, 2)) // = 50
}
