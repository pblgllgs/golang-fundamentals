//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

type Item struct {
	name  string
	state bool
}

func desactivate(item *Item) {
	item.state = false
}

func desactivateAll(items *[4]Item) {
	for i := 0; i < len(*items); i++ {
		items[i].state = false
	}
}

func main() {

	milk := Item{
		name:  "milk",
		state: true}

	cookie := Item{
		name:  "cookie",
		state: false}

	tomato := Item{
		name:  "tomato",
		state: true}

	meat := Item{
		name:  "meat",
		state: true}

	items := [4]Item{milk, cookie, tomato, meat}

	fmt.Println(items)

	desactivate(&items[3])

	fmt.Println(items)

	desactivateAll(&items)

	fmt.Println(items)
}
