package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Contact struct {
	ID    int
	Name  string
	Phone string
}

var contacts []Contact
var lastID int

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
			ReadContacts()

		case "3":
			updateContacts(scanner)

		case "4":
			deleteContact(scanner)

		case "5":
			fmt.Println("exiting...")
			return

		default:
			fmt.Println("invalid option")
		}

	}
}

func createContact(scanner *bufio.Scanner) {
	lastID++
	fmt.Println("enter name: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Println("Enter phone number: ")
	scanner.Scan()
	phone := scanner.Text()

	contact := Contact{ID: lastID, Name: name, Phone: phone}
	contacts = append(contacts, contact)

	fmt.Println("contact added successfully !")
}

func ReadContacts() {
	fmt.Println("\n-- Contact List --")
	for _, contact := range contacts {
		fmt.Printf("ID: %d \n Name: %s \n Phone: %s \n", contact.ID, contact.Name, contact.Phone)
	}
}

func updateContacts(scanner *bufio.Scanner) {
	fmt.Print("Enter id will update: ")
	scanner.Scan()
	id, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("invalid id")
		return
	}

	for i, contact := range contacts {
		if contact.ID == id {
			fmt.Print("NewName: ")
			scanner.Scan()
			name := scanner.Text()

			fmt.Print("New Phone: ")
			scanner.Scan()
			phone := scanner.Text()

			contacts[i].Name = name
			contacts[i].Phone = phone

			fmt.Println("update successfully")
			return

		}
	}

	fmt.Println("contact not found")
}

func deleteContact(scanner *bufio.Scanner) {
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
			fmt.Println("Contact deleted successfully!")
			return
		}
	}

	fmt.Println("Contact not found.")
}
