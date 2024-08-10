package main

import "fmt"

func main() {
	employees := map[int]string{
		1: "kko",
		2: "Ram",
	}
	// map bisa mengembalikan 2 value yaitu value, dan boolean
	value1, ok := employees[1] // ok=> boolean
	fmt.Println(value1)
	fmt.Println(ok)

	if !ok {
		fmt.Println("value not found in key 1")
	} else {
		fmt.Println("value 1: ", value1)
	}

	value2, ok := employees[10]
	if !ok {
		fmt.Println("value not found key 10")
	} else {
		fmt.Println("laue for key 10: ", value2)
	}
}
