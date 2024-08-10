package main

import (
	"fmt"
)

func main() {
	for start := byte('a'); start <= byte('z'); start++ {
		fmt.Print(string(start) + " ")
	}
	fmt.Println()
}
