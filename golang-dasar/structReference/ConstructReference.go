package main

import "fmt"

type Employee struct {
	name string
	id   int
}

func main() {
	emp1 := &Employee{}
	emp1.name = "koko"
	emp1.id = 12

	emp2 := &Employee{"JOni", 13}

	fmt.Println("Name: ", emp1.name)
	fmt.Println("Id: ", emp1.id)
	fmt.Println("")

	fmt.Println("Name: ", emp2.name)
	fmt.Println("Id: ", emp2.id)

}
