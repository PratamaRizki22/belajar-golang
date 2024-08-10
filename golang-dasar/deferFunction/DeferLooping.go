package main

import "fmt"

func main() {
	fmt.Println("start")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i) //dijalankan terakhir dan terbalik
	}

	fmt.Println("finish")
}
