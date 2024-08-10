package main

import "fmt"

func readData2(message string) string {
	fmt.Println(message)
	var data string

	fmt.Scanln(&data)

	return data
}

func main() {
	userName := readData2("enter name: ")
	city := readData2("enter city: ")
	country := readData2("enter country: ")

	fmt.Println("\nUser Entered\n")
	fmt.Println("User name: ", userName)
	fmt.Println("city: ", city)
	fmt.Println("country: ", country)
}
