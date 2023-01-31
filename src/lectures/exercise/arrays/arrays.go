//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name  string
	price int
}

func lastProduct(arr []Product) int {
	return count(arr) - 1
}

func count(arr []Product) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		if arr[i].name != "" {
			sum += 1
		}
	}
	return sum
}

func totalCost(arr []Product) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		if arr[i].name != "" {
			sum += arr[i].price
		}
	}
	return sum
}

func info(arr []Product) {
	fmt.Println("the last product is", arr[lastProduct(arr[:])].name)
	fmt.Println("the total of products is", count(arr[:]))
	fmt.Println("the total cost of products is", totalCost(arr[:]))
}

func main() {
	products := [4]Product{
		{"Bnana", 1},
		{"Meat", 6},
		{"Salad", 3},
	}
	info(products[:])
	products[3] = Product{"Brad", 4}
	info(products[:])
}
