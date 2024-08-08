// mirip class di OOP
package main

import "fmt"

type Employee struct {
	name string
	age  int
}

func main() {
	emp1 := Employee{"koko", 21}
	emp2 := Employee{}
	emp2.name = "nono"
	emp2.age = 31

	fmt.Println(emp1.name)
	fmt.Println(emp1.age)

	fmt.Println()
	fmt.Println(emp2.name)
	fmt.Println(emp2.age)
}
