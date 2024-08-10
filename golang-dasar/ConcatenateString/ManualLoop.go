package main

import "fmt"

func main() {
	data := []string{"Hello", "world", "How", "Are", "You"}
	result := ""
	for _, str := range data {
		result = result + str
	}
	fmt.Println(result)
}
