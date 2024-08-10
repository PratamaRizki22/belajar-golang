package main

import (
	"bytes"
	"fmt"
)

func main() {
	data := []string{"Hello", "world", "How", "Are", "You"}
	var buffer bytes.Buffer
	for _, str := range data {
		buffer.WriteString(str)
	}
	fmt.Println(buffer.String())
}
