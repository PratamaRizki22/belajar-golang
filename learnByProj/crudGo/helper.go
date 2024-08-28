package crudgo

import (
	"encoding/json"
	"fmt"
	"os"
)

const fileName = "contacts.json"

func saveContact(contact Contact) {
	contacts := loadContacts()
	contacts = append(contacts, contact)
	saveContactsToFile(contacts)
}

func saveContactsToFile(contacts []Contact) {
	data, err := json.Marshal(contacts)
	if err != nil {
		fmt.Println("Error saving contacts:", err)
		return
	}

	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func loadContacts() []Contact {
	data, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []Contact{}
		}
		fmt.Println("Error loading contacts:", err)
		return []Contact{}
	}

	var contacts []Contact
	err = json.Unmarshal(data, &contacts)
	if err != nil {
		fmt.Println("Error parsing contacts:", err)
		return []Contact{}
	}

	return contacts
}

func getLastID() int {
	contacts := loadContacts()
	lastID := 0
	for _, contact := range contacts {
		if contact.ID > lastID {
			lastID = contact.ID
		}
	}
	return lastID
}
