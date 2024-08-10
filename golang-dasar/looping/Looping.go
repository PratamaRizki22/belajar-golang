package main

import "fmt"

func main() {
	fmt.Println("basic for loop")
	basicloop()
	fmt.Println("")

	fmt.Println("while loop")
	whileLoop()
	fmt.Println("")

	fmt.Println("Infinite loop/ do while loop")
	doWhileLoop()
	fmt.Println("")

	fmt.Println("range loop slice")
	rangeloopslices()
	fmt.Println("")

	fmt.Println("range loop array")
	rangelooparray()
	fmt.Println("")

	fmt.Println("range loop map")
	rangeloopmap()
	fmt.Println("")

	fmt.Println("range loop string")
	rangeloopstring()
	fmt.Println("")

	fmt.Println("nested loop")
	nestedloop()
	fmt.Println("")
}

func basicloop() {
	for i := 0; i < 5; i++ {
		fmt.Println(i, ": basic for loop")
	}
}

func whileLoop() {
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
}

func doWhileLoop() {
	i := 0
	for {
		fmt.Println(i)
		i++
		if i >= 5 {
			break
		}
	}
}

func rangeloopslices() {
	number := []int{1, 2, 3, 4, 5, 6}

	for index, value := range number {
		fmt.Println(index, ": ", value)
	}
}

func rangelooparray() {
	number := [5]int{1, 2, 3, 4, 5}
	for index, value := range number {
		fmt.Println(index, ": ", value)
	}
}

func rangeloopmap() {
	person := map[string]string{
		"Name":   "joko",
		"Age":    "12",
		"status": "mommy",
	}
	for key, value := range person {
		fmt.Println(key, ":", value)
	}
}

func rangeloopstring() {
	text := "hello"

	for i, runeValue := range text {
		fmt.Println(i, ":", runeValue)
	}
}

func nestedloop() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("i: %d, j: %d\n", i, j)
		}
	}
}
