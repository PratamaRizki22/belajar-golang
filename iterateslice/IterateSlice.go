package main

import "fmt"

func main() {
	var vowels = []string{"a", "i", "u", "e", "o"}

	for i := range vowels {
		fmt.Println(vowels[i])
	}

	for i, v := range vowels {
		fmt.Printf("index %v, value %v\n", i, v)
	}
}

//kkoko tampan