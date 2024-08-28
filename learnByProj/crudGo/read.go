package crudgo

import "fmt"

func ReadContacts() {
	contacts := loadContacts()
	fmt.Println("\n--- Contact List ---")
	for _, contact := range contacts {
		fmt.Printf("ID: %d, Name: %s, Phone: %s\n", contact.ID, contact.Name, contact.Phone)
	}
}
