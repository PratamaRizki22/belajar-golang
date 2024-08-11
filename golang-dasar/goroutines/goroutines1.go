package main

import (
	"fmt"
)

func printUpperAlphabets() {
	for alphabet := byte('A'); alphabet <= byte('Z'); alphabet++ {
		fmt.Print(string(alphabet))
	}
	fmt.Println()
}

func printLowerAlphabets() {
	for alphabet := byte('a'); alphabet <= byte('z'); alphabet++ {
		fmt.Print(string(alphabet), " ")
	}
	fmt.Println()
}

func main() {
	printUpperAlphabets()
	printLowerAlphabets()
}
