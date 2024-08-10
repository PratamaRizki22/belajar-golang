package main

import (
	"fmt"
)

func main() {
	msg := "Hello World"
	msgBytes := []byte(msg)
	fmt.Println("msg : ", msg)
	fmt.Println("msgBytes : ", msgBytes)
}
