package main

import "fmt"

type Employee2 struct {
	name    string
	id      int
	hobbies []string
}

func (emp *Employee2) printEmployeeInfo() {
	fmt.Println("name : ", emp.name, ", id : ", emp.id, ", hobbies : ", emp.hobbies)
}

func getEmployee(name string, id int, hobbies []string) (*Employee2, error) {
	if name == "" {
		return nil, fmt.Errorf("Name can't be nil")
	}
	if id <= 0 {
		return nil, fmt.Errorf("id can't be zero")
	}
	emp := Employee2{}
	emp.name = name
	emp.id = id
	emp.hobbies = hobbies
	return &emp, nil
}

func main() {
	emp1, err1 := getEmployee("JJTam", 1, []string{"Cricket", "Football"})
	if err1 != nil {
		fmt.Println("Error Occurred while creating employee")
		fmt.Println(err1)
	} else {
		emp1.printEmployeeInfo()
	}
	emp2, err2 := getEmployee("JJTam", -10, []string{"Cricket", "Football"})
	if err2 != nil {
		fmt.Println("Error Occurred while creating employee")
		fmt.Println(err2)
	} else {
		emp2.printEmployeeInfo()
	}
}
