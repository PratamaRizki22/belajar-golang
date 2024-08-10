package main

import "fmt"

type Employee2 struct {
	Name string
	ID   int8
}

func main() {
	i := 10
	j := 10.11
	k := "JJTam"
	emp := Employee{}
	fmt.Println("Type of i : ", fmt.Sprintf("%T", i))
	fmt.Println("Type of j : ", fmt.Sprintf("%T", j))
	fmt.Println("Type of k : ", fmt.Sprintf("%T", k))
	fmt.Println("Type of emp : ", fmt.Sprintf("%T", emp))
}
