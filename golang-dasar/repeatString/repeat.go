package main

import (
	"fmt"
	"strings"
)

func main() {
	msg1 := strings.Repeat("*", 50)
	msg2 := strings.Repeat("-", 50)

	fmt.Println(msg1)
	fmt.Println(msg2)
}
