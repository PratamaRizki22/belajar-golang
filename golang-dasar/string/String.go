package main

import "fmt"

// string immutable =tidak bisa diubah
func main() {
	var msg string = "Hello World"
	msg2 := ", kamu keren"

	fmt.Printf("msg : %v, type : %T\n", msg, msg)
	fmt.Printf("Length of msg is %v\n", len(msg))

	for i := 0; i < len(msg); i++ {
		fmt.Printf("msg[%v]=%v\n", i, msg[i])
	}

	result := msg + msg2

	fmt.Println(result)
}
