package main

import "fmt"

func main() {
	i := int8(10)
	j := float32(10.23)
	k := bool(true)
	l := "Hello world"

	fmt.Printf("i is of type %T\n", i)
	fmt.Printf("j is of type %T\n", j)
	fmt.Printf("k is of type %T\n", k)
	fmt.Printf("l is of type %T\n", l)
}
