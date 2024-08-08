package main

import "fmt"

func main() {
	result := add(1, 2, 3)

	fmt.Println(result)
}

// hanya boleh ketika semua parameter tipenya sama
func add(x, y, z int8) int8 {
	return x + y + z
}
