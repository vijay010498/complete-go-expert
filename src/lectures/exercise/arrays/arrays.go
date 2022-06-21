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

func printInformation(products [4]Product) {
	var cost, totalItems int

	for i := 0; i < len(products); i++ {
		item := products[i]
		cost += item.price

		if item.name != "" {
			totalItems += 1
		}
		fmt.Println("Last item:", products[totalItems-1])
		fmt.Println("Total Items", totalItems)
		fmt.Println("Total Cost", cost)
	}
}

func main() {
	products := [4]Product{
		{name: "1", price: 10},
		{name: "2", price: 10},
		{name: "3", price: 10},
	}
	printInformation(products)
}
