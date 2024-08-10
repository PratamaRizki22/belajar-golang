// dalam bahasa go sebenarnya tidak ada konsep consturtor namun bisa dibuat sendiri dengan function

package main

import (
	"fmt"
	"strings"
)

type Employee struct {
	name    string
	id      int8
	hobbies []string
}

func getEmployess(name string, id int8, hobbies []string) *Employee {
	emp := &Employee{}
	emp.name = name
	emp.id = id
	emp.hobbies = hobbies

	return emp
}

func main() {
	emp := getEmployess("joko", 12, []string{"makan", "tidur", "joget", "makan", "tidur", "makan", "tidur"})
	hobbieStr := strings.Join(emp.hobbies, ", ")
	fmt.Printf(`
nama: %s
id: %d
hobbie: %v 
`, emp.name, emp.id, hobbieStr)
}
