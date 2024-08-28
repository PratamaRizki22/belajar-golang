package crudgo

import (
	"bufio"
	"fmt"
)

type Contact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

var lastID int

func CreateContact(scanner *bufio.Scanner) {
	lastID = getLastID() + 1
	fmt.Print("Enter Name: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Print("Enter Phone: ")
	scanner.Scan()
	phone := scanner.Text()

	contact := Contact{ID: lastID, Name: name, Phone: phone}
	saveContact(contact)

	fmt.Println("Contact added successfully!")
}
