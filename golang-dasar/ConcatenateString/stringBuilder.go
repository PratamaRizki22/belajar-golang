package main

//rekomnded
import (
	"fmt"
	"strings"
)

func main() {
	data := []string{"Hello", "world", "How", "Are", "You"}
	var result strings.Builder
	for _, str := range data {
		result.WriteString(str)
	}
	fmt.Println(result.String())
}
