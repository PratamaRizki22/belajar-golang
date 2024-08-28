package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// Struct untuk menyimpan informasi kontak
type Contact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

var lastID int

const fileName = "contacts.json"

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n--- Contact Management CLI ---")
		fmt.Println("1. Create Contact")
		fmt.Println("2. Read Contacts")
		fmt.Println("3. Update Contact")
		fmt.Println("4. Delete Contact")
		fmt.Println("5. Exit")
		fmt.Print("Select an option: ")

		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			createContact(scanner)
		case "2":
			readContacts()
		case "3":
			updateContact(scanner)
		case "4":
			deleteContact(scanner)
		case "5":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please select again.")
		}
	}
}

func createContact(scanner *bufio.Scanner) {
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

func readContacts() {
	contacts := loadContacts()
	fmt.Println("\n--- Contact List ---")
	for _, contact := range contacts {
		fmt.Printf("ID: %d, Name: %s, Phone: %s\n", contact.ID, contact.Name, contact.Phone)
	}
}

func updateContact(scanner *bufio.Scanner) {
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

func deleteContact(scanner *bufio.Scanner) {
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

	err = ioutil.WriteFile(fileName, data, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func loadContacts() []Contact {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []Contact{} // Return empty slice if file does not exist
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
