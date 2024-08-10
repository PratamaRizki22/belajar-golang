package main

import (
	"crypto/rand"
	"fmt"
)

func GenerateRandomString() *string {
	tempBytes := make([]byte, 16)
	_, err := rand.Read(tempBytes)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil
	}
	uuid := fmt.Sprintf("%X-%X-%X-%X-%X", tempBytes[0:4], tempBytes[4:6],
		tempBytes[6:8], tempBytes[8:10], tempBytes[10:])
	return &uuid
}
func main() {
	randomString := *GenerateRandomString()
	fmt.Println(randomString)
}
