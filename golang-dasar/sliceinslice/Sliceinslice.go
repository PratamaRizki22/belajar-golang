package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	subslice := slice[3:6]

	fmt.Println("slice: ", slice)
	fmt.Println("subslice: ", subslice)
}
