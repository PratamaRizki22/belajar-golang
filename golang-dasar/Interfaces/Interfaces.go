package main

import "fmt"

type Employee struct {
	name string
	id   int8
}

func (emp *Employee) ChangeId(newId int8) int8 {
	emp.id = newId
	return emp.id
}

func (emp *Employee) aboutMe() {
	fmt.Println("name: ", emp.name, ", id: ", emp.id)
}

type Detail interface {
	ChangeId(NewId int8) int8
	aboutMe()
}

func printMyDetail(obj Detail) {
	obj.aboutMe()
}

func main() {
	emp1 := &Employee{"koko", 1}
	printMyDetail(emp1)
	emp1.ChangeId(12)
	printMyDetail(emp1)

}
