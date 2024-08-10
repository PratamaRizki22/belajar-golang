package main

import (
	"fmt"
	"strings"
)

type Employee2 struct {
	name string
	id   int8
}

func (emp Employee2) printInfo() (string, int8) {
	return emp.name, emp.id
}

func (emp *Employee2) capitalizeName() {
	emp.name = strings.ToUpper(emp.name)
}

func main() {
	// emp1 := &Employee2{}
	emp1 := new(Employee2)

	emp1.name = "Jojo"
	emp1.id = 12

	name, id := emp1.printInfo()

	fmt.Println(name)
	fmt.Println(id)

	emp1.capitalizeName()
	name, id = emp1.printInfo()

	fmt.Println(name)
	fmt.Println(id)

}
