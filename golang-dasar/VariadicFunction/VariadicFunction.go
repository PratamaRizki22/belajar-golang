package main

import "fmt"

func main() {
	sayHello("koko")
	sayHello("hoodo")
	sayHello("Nmaya")
}

func sayHello(names ...string) {
	// fmt.Println(names, "\n") //redudan list tidak boleh
	for _, name := range names {
		fmt.Println("Hello ", name)
	}
}
