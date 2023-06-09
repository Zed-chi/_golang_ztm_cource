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

import (
	"fmt"
)

type Product struct {
	name  string
	price float32
}

func sum(qwe [4]Product) float32 {
	var ttl float32 = 0
	for i := 0; i < len(qwe); i++ {
		ttl += (qwe)[i].price
	}
	return ttl
}
func main() {
	var list [4]Product

	list[0] = Product{"qwe", 123}
	list[1] = Product{"asd", 123}
	list[2] = Product{"zxc", 123}

	fmt.Println(list[0])
	fmt.Println(len(list))
	fmt.Println(sum(list))

	list[3] = Product{"dfg", 123}
	fmt.Println(list[0])
	fmt.Println(len(list))
	fmt.Println(sum(list))
}
