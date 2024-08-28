package crudgo

import (
	"bufio"
	"fmt"
	"strconv"
)

func DeleteContact(scanner *bufio.Scanner) {
	contacts := loadContacts()
	fmt.Print("Enter the ID of the contact to delete: ")
	scanner.Scan()
	id, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	for i, contact := range contacts {
		if contact.ID == id {
			contacts = append(contacts[:i], contacts[i+1:]...)
			saveContactsToFile(contacts)
			fmt.Println("Contact deleted successfully!")
			return
		}
	}

	fmt.Println("Contact not found.")
}
