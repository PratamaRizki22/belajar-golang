package main

import (
	"fmt"
	"reflect"
)

type Employee struct {
	Name string
	ID   int8
}

func main() {
	i := 10
	j := 10.11
	k := "koko"
	emp := Employee{}

	fmt.Println("Type of i : ", reflect.TypeOf(i))
	fmt.Println("Type of j : ", reflect.TypeOf(j))
	fmt.Println("Type of k : ", reflect.TypeOf(k))
	fmt.Println("Type of emp : ", reflect.TypeOf(emp))
}
