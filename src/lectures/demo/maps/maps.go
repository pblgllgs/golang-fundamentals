package main

import "fmt"

func main() {
	shoppingList := make(map[string]int)
	shoppingList["eggs"] = 11
	shoppingList["milk"] = 1
	shoppingList["bread"] += 1

	shoppingList["eggs"] += 1

	fmt.Println(shoppingList)

	delete(shoppingList, "milk")

	fmt.Println("Milk deleted, ner list: ", shoppingList)

	fmt.Println("need", shoppingList["eggs"], "eggs")

	cereal, found := shoppingList["cereal"]
	fmt.Println("Need cereal?")
	if !found {
		fmt.Println("not cereal")
	} else {
		fmt.Println("yes!", cereal, "boxes")
	}

	totalCount := 0
	for _, amount := range shoppingList {
		totalCount += amount
	}

	fmt.Println("count", totalCount)
}
