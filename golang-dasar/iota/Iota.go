package main

import (
	"fmt"
	"unsafe"
)

const (
	block1first = iota
	block1second
)

const (
	first int8 = iota
	second
	third
	fourth
)

//  << shift digunakan untuk geser bit
//  misalnya x= 00000011 adalah 3 dalam decimal
// x << 1 menjadi 00000110 adalah 6 dalam decimal
// x << 2 menjadi 0001100 adalah 12 dalam decimal

// iota datanya tetap yaitu integer walau diubah apapun tetap integer
const (
	_  int8 = iota             // ignore first value by assigning to blank identifier
	KB      = 1 << (10 * iota) // 1 KB = 1024 bytes
	MB                         // 1 MB = 1024 * 1024 bytes
	GB                         // 1 GB = 1024 * 1024 * 1024 bytes
	TB                         // 1 TB = 1024 * 1024 * 1024 * 1024 bytes
)

func main() {
	fmt.Printf("block1first: %v \n", block1first)
	fmt.Printf("block1seocnd: %v \n", block1second)
	fmt.Printf("first: %v \n", first)
	fmt.Printf("second: %v \n", second)
	fmt.Printf("third: %v \n", third)
	fmt.Printf("fourth: %v \n", fourth)

	fmt.Printf("\nMemory Sizes using iota:\n")
	fmt.Printf("1 KB: %d bytes, type=%T, size=%d bytes\n", KB, KB, unsafe.Sizeof(KB))
	fmt.Printf("1 MB: %d bytes, type=%T, size=%d bytes\n", MB, MB, unsafe.Sizeof(MB))
	fmt.Printf("1 GB: %d bytes, type=%T, size=%d bytes\n", GB, GB, unsafe.Sizeof(GB))
	fmt.Printf("1 TB: %d bytes, type=%T, size=%d bytes\n", TB, TB, unsafe.Sizeof(TB))

	// iota()
}
