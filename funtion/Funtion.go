package main

import "fmt"

func main() {
	sayhello()
	person := sayname("rizki")
	fmt.Println(person)
}

func sayhello() {
	fmt.Println("hello")
}

func sayname(name string) string {
	return ("hello " + name)
}
