package main

import "fmt"

func main() {
	var i int
	var j float32 = 2.34
	i = int(j)

	var k float32
	l := int(2)
	k = float32(l)

	fmt.Printf("i= %v , type=%T \n", i, i)
	fmt.Printf("j= %v, type=%T\n", j, j)

	fmt.Printf("k= %v, type=%T\n", k, k)
	fmt.Printf("l= %v, trpe=%T\n", l, l)
}
