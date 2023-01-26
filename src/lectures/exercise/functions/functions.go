//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import (
	"fmt"
)

func greeting(name string) {
	fmt.Println("Nice to meet you", name)
}

func message(msg string) string {
	return msg
}

func sum(one, two, three int) int {
	return one + two + three
}

func randomNumber(num int) int {
	return num
}

func randomNumberTwo(num1, num2 int) (int, int) {
	return num1, num2
}

func unionPrint(num1, num2, num3 int) {
	uno := randomNumber(num1)
	dos, tres := randomNumberTwo(num2, num3)
	result := uno + dos + tres
	fmt.Println(result)
}

func main() {
	greeting("Pablo")
	fmt.Println(message("Holaaa in golang"))
	suma := sum(1, 2, 3)
	fmt.Println(suma)
	valor1 := randomNumber(5)
	fmt.Println(valor1)
	valor2, valor3 := randomNumberTwo(5, 6)
	fmt.Println(valor2, valor3)
	unionPrint(1, 2, 3)
}
