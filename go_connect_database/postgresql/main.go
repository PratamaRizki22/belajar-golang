package main

import (
	"fmt"
	"log"

	"postgres_connect/db"    // Mengimpor package `db`
	"postgres_connect/model" // Mengimpor package `model`
)

func main() {
	// Koneksi ke database
	dbConn, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	} else {
		log.Println("Successfully connected to the database!")

	}

	defer dbConn.Close()

	// Buat tabel pengguna
	err = model.CreateUserTable(dbConn)
	if err != nil {
		log.Fatal("Error creating table:", err)
	} else {
		fmt.Println("Table created successfully!")
	}
}
