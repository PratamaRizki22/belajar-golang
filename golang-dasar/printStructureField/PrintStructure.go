package main

import "fmt"

type employee struct {
	FirstName string
	LastName  string
	id        int8
}

func main() {
	emp1 := employee{"joko", "susilo", 12}

	fmt.Printf("%+v\n", emp1)
}
