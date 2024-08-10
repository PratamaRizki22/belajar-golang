package main

import "fmt"

func main() {
	// Contoh Shift Kiri (<<)
	var a int = 1
	var shiftLeft1 int = a << 1 // Geser kiri 1 posisi
	var shiftLeft2 int = a << 2 // Geser kiri 2 posisi

	// Visualisasi Shift Kiri
	fmt.Printf("Sebelum shift kiri: %08b (decimal: %d)\n", a, a)
	fmt.Printf("Setelah shift kiri 1 posisi: %08b (decimal: %d)\n", shiftLeft1, shiftLeft1)
	fmt.Printf("Setelah shift kiri 2 posisi: %08b (decimal: %d)\n\n", shiftLeft2, shiftLeft2)

	// Contoh Shift Kanan (>>)
	var b int = 8
	var shiftRight1 int = b >> 1 // Geser kanan 1 posisi
	var shiftRight2 int = b >> 2 // Geser kanan 2 posisi

	// Visualisasi Shift Kanan
	fmt.Printf("Sebelum shift kanan: %08b (decimal: %d)\n", b, b)
	fmt.Printf("Setelah shift kanan 1 posisi: %08b (decimal: %d)\n", shiftRight1, shiftRight1)
	fmt.Printf("Setelah shift kanan 2 posisi: %08b (decimal: %d)\n\n", shiftRight2, shiftRight2)

	// Menggunakan iota untuk Mendefinisikan Ukuran Memori
	const (
		_  = iota
		KB = 1 << (10 * iota) // 1 KB = 1024 bytes
		MB                    // 1 MB = 1024 * 1024 bytes
		GB                    // 1 GB = 1024 * 1024 * 1024 bytes
		TB                    // 1 TB = 1024 * 1024 * 1024 * 1024 bytes
	)

	// Visualisasi Ukuran Memori
	fmt.Printf("1 KB: %d bytes\n", KB) // Output: 1 KB: 1024 bytes
	fmt.Printf("1 MB: %d bytes\n", MB) // Output: 1 MB: 1048576 bytes
	fmt.Printf("1 GB: %d bytes\n", GB) // Output: 1 GB: 1073741824 bytes
	fmt.Printf("1 TB: %d bytes\n", TB) // Output: 1 TB: 1099511627776 bytes
}
