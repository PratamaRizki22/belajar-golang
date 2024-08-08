package main

import "fmt"

func main() {
	a := 10
	switch {
	case a > 10:
		fmt.Println("lebih dari 10")
	case a <= 10:
		fmt.Println("kurang dari 10")
	}

}
