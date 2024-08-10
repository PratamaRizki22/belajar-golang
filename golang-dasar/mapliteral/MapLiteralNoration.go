package main

import "fmt"

func main() {
	person := map[string]string{
		"Name":   "joko",
		"Age":    "12",
		"status": "mommy",
	}
	fmt.Println(person)
	fmt.Println()

	for key, value := range person {
		fmt.Println(key, ":", value)
	}
}
