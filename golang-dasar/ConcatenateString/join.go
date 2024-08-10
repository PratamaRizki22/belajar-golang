package main

import (
	"fmt"
	"strings"
)

func main() {
	data := []string{"Hello", "world", "How", "Are", "You"}
	result := strings.Join(data, "")
	fmt.Println(result)
}
