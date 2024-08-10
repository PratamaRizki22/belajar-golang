package main

import "fmt"

func main() {
	vowels := []string{"a", "b", "c"}

	vowels = append(vowels, "d")
	vowels = append(vowels, "e")
	for index, value := range vowels {
		fmt.Println(index, ": ", value)

	}
}
