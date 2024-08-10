package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// string adalah slice of bytes yang sifatnya immutable
// sedangkan slice of bytes sendiri di golang sifatnya muttable

func main() {
	s1 := "hello"
	b1 := []byte(s1) // Menghasilkan slice of bytes: ['h', 'e', 'l', 'l', 'o']

	fmt.Println(string(b1))
	fmt.Println(b1)

	b2 := []byte{'h', 'e', 'l', 'l', 'o'}
	s2 := string(b2) // Menghasilkan string "hello"
	fmt.Println(s2)
	fmt.Println([]byte(s2))

	fmt.Println("pembuktian")
	pembuktianbuatmanual()

}

func pembuktianbuatmanual() {
	// Buat array byte secara manual
	byteArray := []byte{104, 101, 108, 108, 111} // representasi "hello"

	// Buat struct string secara manual menggunakan unsafe
	var str string
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	stringHeader.Data = uintptr(unsafe.Pointer(&byteArray[0]))
	stringHeader.Len = len(byteArray)

	fmt.Println("before: ", byteArray)
	// Tampilkan string yang dibuat
	fmt.Println(str) // Output: "hello"
}
