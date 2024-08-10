package main

import "fmt"

type Employee struct {
	name string
	id   int8
}

func (emp Employee) printEmployeeInfo() {
	fmt.Println("name: ", emp.name, ", id: ", emp.id)
}

func main() {
	emp1 := new(Employee)
	emp1.name = "joko"
	emp1.id = 12

	emp1.printEmployeeInfo()
}
