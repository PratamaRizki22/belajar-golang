package crudgo

import (
	"bufio"
	"fmt"
	"strconv"
)

func UpdateContact(scanner *bufio.Scanner) {
	contacts := loadContacts()
	fmt.Print("Enter the ID of the contact to update: ")
	scanner.Scan()
	id, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	for i, contact := range contacts {
		if contact.ID == id {
			fmt.Print("Enter new Name: ")
			scanner.Scan()
			name := scanner.Text()

			fmt.Print("Enter new Phone: ")
			scanner.Scan()
			phone := scanner.Text()

			contacts[i].Name = name
			contacts[i].Phone = phone
			saveContactsToFile(contacts)

			fmt.Println("Contact updated successfully!")
			return
		}
	}

	fmt.Println("Contact not found.")
}
