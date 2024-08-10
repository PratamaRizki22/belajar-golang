package main

import "fmt"

//Fungsi yang sama dari struct yang berbeda bisa dijalankan

// Interface yang mendefinisikan metode getName
type Named interface {
	getName() string
}

// Struct Employee2 yang mengimplementasikan interface Named
type Employee2 struct {
	name string
	id   int
}

func (e Employee2) getName() string {
	return e.name
}

// Struct Manager2 yang juga mengimplementasikan interface Named
type Manager2 struct {
	name string
	id   int
	dept string
}

func (m Manager2) getName() string {
	return m.name
}

// Fungsi printName2 yang sekarang menggunakan interface Named
func printName2(n Named) {
	fmt.Println(n.getName())
}

func main() {
	emp := Employee2{name: "John", id: 123}
	mgr := Manager2{name: "Jane", id: 456, dept: "Sales"}

	// Keduanya dapat digunakan dengan fungsi printName2
	printName2(emp) // Output: John
	printName2(mgr) // Output: Jane
}
