package main

import (
	"fmt"
)

func main() {
	data := []string{"Hello", "world", "How", "Are", "You"}
	result := make([]byte, 100)
	currentLength := 0
	for _, str := range data {
		currentLength += copy(result[currentLength:], []byte(str))
	}
	fmt.Println(string(result))
}
