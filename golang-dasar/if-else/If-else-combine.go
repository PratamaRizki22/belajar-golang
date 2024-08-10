package main

import "fmt"

func main() {
	if b := 12; b == 13 {
		fmt.Println(b)
	} else {
		fmt.Println("not", b)
	}
}
