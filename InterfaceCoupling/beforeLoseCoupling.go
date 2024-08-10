package main

import "fmt"

type Employee struct {
	name string
	id   int
}

func (e Employee) getName() string {
	return e.name
}

type Manager struct {
	name string
	id   int
	dept string
}

func (m Manager) getName() string {
	return m.name
}

func printName(emp Employee) {
	fmt.Println(emp.getName())
}

func main() {
	emp := Employee{name: "John", id: 123}
	printName(emp)

	// Ini tidak akan berfungsi karena fungsi printName hanya menerima Employee, bukan Manager
	// mgr := Manager{name: "Jane", id: 456, dept: "Sales"}
	// printName(mgr)
}
