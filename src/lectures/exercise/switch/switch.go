//--Summary:
//  Create a program to display a classification based on age.
//
//--Requirements:
//* Use a `switch` statement to print the following:
//  - "newborn" when age is 0
//  - "toddler" when age is 1, 2, or 3
//  - "child" when age is 4 through 12
//  - "teenager" when age is 13 through 17
//  - "adult" when age is 18+

package main

import "fmt"

func calc(num int) int {
	return num
}

func main() {

	switch age := calc(18); {
	case age == 0:
		fmt.Println("newborn")
	case age == 1 || age == 2 || age == 3:
		fmt.Println("toddler")
	case age >= 4 && age <= 12:
		fmt.Println("child")
	case age >= 13 && age <= 17:
		fmt.Println("teenager")
	default:
		fmt.Println("adult")
	}

}
