package main

import (
	"bufio"
	"fmt"
	"os"
)

func readData(message string) string {
	fmt.Println(message)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	if scanner.Err() != nil {
		return "Can't read input"
	}
	return scanner.Text()
}

func main() {
	userName := readData("Enter your name: ")
	county := readData("Enter your country")

	fmt.Println("User Entered\n")
	fmt.Println("User Name: ", userName)
	fmt.Println("country: ", county)
}
