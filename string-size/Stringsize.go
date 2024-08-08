package main

import (
	"fmt"
	"unicode/utf8"
	"unsafe"
)

func main() {
	str := "Hello, 世界"
	str2 := "12"

	// Panjang string dalam byte
	lengthInBytes := len(str)
	lengthInBytes2 := len(str2)

	// Panjang string dalam rune (karakter Unicode)
	lengthInRunes := utf8.RuneCountInString(str)
	lengthInRunes2 := utf8.RuneCountInString(str2)

	// Ukuran struktur string (header)
	sizeOfStringHeader := unsafe.Sizeof(str)
	sizeOfStringHeader2 := unsafe.Sizeof(str2)

	fmt.Printf("String: %s\n", str)
	fmt.Printf("Length in bytes: %d\n", lengthInBytes)
	fmt.Printf("Length in characters (runes): %d\n", lengthInRunes)
	fmt.Printf("Size of string header: %d bytes\n", sizeOfStringHeader)
	fmt.Printf("Total size including header and data: %d bytes\n", sizeOfStringHeader+uintptr(lengthInBytes))
	fmt.Println("")

	fmt.Printf("String: %s\n", str2)
	fmt.Printf("Length in bytes: %d\n", lengthInBytes2)
	fmt.Printf("Length in characters (runes): %d\n", lengthInRunes2)
	fmt.Printf("Size of string header: %d bytes\n", sizeOfStringHeader2)
	fmt.Printf("Total size including header and data: %d bytes\n", sizeOfStringHeader2+uintptr(lengthInBytes2))

}
