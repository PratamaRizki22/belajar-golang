package main

import "fmt"

func main() {
	name := new(string)
	*name = "Ram"

	changeMe(name) // & artinya reference
	fmt.Println(*name)
	// Output: koko

	// Penjelasan tambahan:
	// - `&name` memberikan alamat memori dari variabel `name` ke fungsi `changeMe`.
	// - `changeMe` menggunakan pointer ini untuk mengubah nilai asli dari variabel `name`.
}

func changeMe(name *string) {
	*name = "koko" // Mengubah nilai asli yang ditunjuk oleh pointer `name` menjadi "koko"
}
