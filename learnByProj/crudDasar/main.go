package main

import (
	"bufio"
	"fmt"
	"lernGo/crudGo"
	"os"
)

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
			crudgo.CreateContact(scanner)
		case "2":
			crudgo.ReadContacts()
		case "3":
			crudgo.UpdateContact(scanner)
		case "4":
			crudgo.DeleteContact(scanner)
		case "5":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please select again.")
		}
	}
}
