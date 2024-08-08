package main

import "fmt"

func main() {
	switch a := 10; a {
	case 10:
		fmt.Println(a)
	case 20:
		fmt.Println(20)

	default:
		fmt.Println(" default block is executed")
	}

}
