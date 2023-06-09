package main

import "fmt"

func main() {
	var arr1 [3]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{1, 2, 3}
	fmt.Println(arr1, arr2, arr3)
}
