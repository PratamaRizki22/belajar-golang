package main

import (
	"fmt"
	"reflect"
)

type Employee3 struct {
	Name string
	ID   int
}

func main() {
	i := 10
	j := 10.11
	k := "JJTam"
	emp := Employee{}
	fmt.Println("Type of i : ", reflect.ValueOf(i).Kind())
	fmt.Println("Type of j : ", reflect.ValueOf(j).Kind())
	fmt.Println("Type of k : ", reflect.ValueOf(k).Kind())
	fmt.Println("Type of emp : ", reflect.ValueOf(emp).Kind())
}
