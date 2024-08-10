package main

import (
	"fmt"
)

func main() {
	msg1 := "Hello"
	msg2 := "World"
	result := fmt.Sprint(msg1, msg2)
	fmt.Println(result)
}
