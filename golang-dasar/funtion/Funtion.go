package main

import "fmt"

func main() {
	sayhello()
	hello, nanya := sayname("rizki")
	fmt.Println(hello)
	fmt.Println(nanya)

	add, sub, mul, div := aritmatic(12, 3)

	fmt.Println(add)
	fmt.Println(sub)
	fmt.Println(mul)
	fmt.Println(div)
}

func sayhello() {
	fmt.Println("hello")
}

// multiple say name
func sayname(name string) (string, string) {
	return ("hello " + name), ("kamu " + name + "?")
}

// multiple return named
func aritmatic(a int, b int) (sum int, sub int, mul int, div int) {
	sum = a + b
	sub = a - b
	mul = a * b
	div = a / b

	return
}
