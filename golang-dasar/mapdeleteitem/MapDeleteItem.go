package main

import "fmt"

func main() {
	employees := map[int]string{
		1: "koko",
		2: "ram",
		3: "koko",
	}

	fmt.Println("employees: ", employees)

	delete(employees, 2) //hapus berdasarkan key
	fmt.Println("employees: ", employees)
}
