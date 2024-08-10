package main

import "fmt"

func main() {
	arr := []byte{97, 98, 99, 100, 101, 102, 103}

	msg := string(arr)

	fmt.Println("msg: ", msg) //a,b,c,d,e,f,g
}
