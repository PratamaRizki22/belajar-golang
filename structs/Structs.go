// mirip class di OOP
package main

import "fmt"

type Address struct {
	City, State string
}
type Employee struct {
	name    string
	age     int
	Address //embeded struct
}

func main() {
	emp1 := Employee{"koko", 21, Address{"Blora", "Jawa tengah"}}
	emp2 := Employee{}
	emp3 := new(Employee)

	emp2.name = "nono"
	emp2.age = 31
	emp2.Address.City = "jatim"
	emp2.Address.State = "ponorogo"

	emp3.name = "masha"
	emp3.age = 33
	emp3.Address.City = "depok"
	emp3.Address.State = "kebangsaan"

	fmt.Println()
	fmt.Println(emp1.name)
	fmt.Println(emp1.age)
	fmt.Println(emp1.City)
	fmt.Println(emp1.State)

	fmt.Println()
	fmt.Println(emp2.name)
	fmt.Println(emp2.age)
	fmt.Println(emp2.City)
	fmt.Println(emp2.State)

	fmt.Println()
	fmt.Println(emp3.name)
	fmt.Println(emp3.age)
	fmt.Println(emp3.City)
	fmt.Println(emp3.State)
}
