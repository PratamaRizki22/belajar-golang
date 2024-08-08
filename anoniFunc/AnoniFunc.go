package main

import "fmt"

func main() {
	add := func(a int, b int) int {
		return a + b
	}
	result := add(10, 20)
	fmt.Println(result)
}
