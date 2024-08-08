package main

import "fmt"

func main() {

	var name string
	var id int8

	fmt.Println("masukkan nama: ")
	fmt.Scanln(&name)

	fmt.Println("masukkan id: ")
	fmt.Scanln(&id)

	fmt.Println("name: ", name, ", id: ", id)
}
