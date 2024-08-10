package main

import "fmt"

const (
	first  = 1 << iota // 00000001
	second             // 00000010
	third              // 00000100
	fourth             // 00001000
	five               // 00010000
	six                // 00100000
)

func main() {
	fmt.Println("first : ", first)
	fmt.Println("second : ", second)
	fmt.Println("third : ", third)
	fmt.Println("fourth : ", fourth)
	fmt.Println("five : ", five)
	fmt.Println("six : ", six)
}
