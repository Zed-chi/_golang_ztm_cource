package main

import "fmt"

func main() {

	type Sample struct {
		qwe  string
		a, b int
	}

	//excplicit
	data := Sample{"qwe", 1, 2}

	//anon
	data2 := struct {
		a string
		b int
	}{"qwe", 1}

	fmt.Print(data, data2)
}
