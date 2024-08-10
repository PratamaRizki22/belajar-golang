package main

import "fmt"

// * dibaca pointer ke
func main() {
	var intPtr *int //mengimpan alamat memory yang tipenya integer
	// default = nil
	var a int = 10

	intPtr = &a

	*intPtr = 12

	fmt.Println(a)
	fmt.Println("\na is stored at address : ", intPtr)
	fmt.Println("Value of a is : ", *intPtr)
}
