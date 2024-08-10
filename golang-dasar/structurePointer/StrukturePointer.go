package main

import "fmt"

type employee struct {
	name string
	id   int
}

func main() {
	emp1 := employee{"koko", 1}
	fmt.Println("name: ", emp1.name, ", id: ", emp1.id)

	var empPtr *employee = &emp1
	// empPtr := &emp1

	empPtr.name = "joko"
	empPtr.id = 12

	fmt.Println("\n name: ", emp1.name, ", id: ", emp1.id)

}
