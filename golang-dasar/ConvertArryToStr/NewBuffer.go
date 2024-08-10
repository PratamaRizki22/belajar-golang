package main

import (
	"bytes"
	"fmt"
)

func main() {
	byteArray := []byte{97, 98, 99, 100, 101, 102, 103}
	msg := bytes.NewBuffer(byteArray).String()
	fmt.Println("msg : ", msg)
}
