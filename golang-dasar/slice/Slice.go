package main

import (
	"fmt"
	"unsafe"
)

// bersifat bisa diubah
func UpdateData(data []int) {
	data[0] = 10
	data[1] = 12
	data[2] = 14
}

// lebih baik pakai make jika tahu karena efisien memory
// make cocok untuk data besar
// langsung juga bagus
// tidak disarankan pakai []int{} kalau tidak tau ukuran
func main() {
	data := make([]int, 0, 9) //sudah ada index 0
	// data := []int{} // masih kosonng
	data2 := []int{10, 12, 14, 4, 5, 6, 7, 11, 3}

	fmt.Println("Length:", len(data))
	fmt.Println("Capacity:", cap(data))
	fmt.Println("Data:", data)

	data = append(data, 1)
	data = append(data, 2)
	data = append(data, 3)
	data = append(data, 4)
	data = append(data, 5)
	data = append(data, 6)
	data = append(data, 7)
	data = append(data, 11)
	data = append(data, 3)

	fmt.Println("after append:", data)

	UpdateData(data)
	fmt.Println("after update:", data)

	// Menghitung ukuran slice dalam byte
	headerSize := unsafe.Sizeof(data)
	elementSize := unsafe.Sizeof(data[0])
	dataLength := len(data)
	dataCapacity := cap(data)
	dataSize := uintptr(dataLength) * elementSize
	totalCapacitySize := uintptr(dataCapacity) * elementSize
	totalSizeInBytes := headerSize + dataSize
	totalCapacitySizeInBytes := headerSize + totalCapacitySize

	fmt.Printf("Ukuran header slice: %d byte\n", headerSize)
	fmt.Printf("Ukuran elemen slice: %d byte\n", elementSize)
	fmt.Printf("Panjang slice: %d\n", dataLength)
	fmt.Printf("Kapasitas slice: %d\n", dataCapacity)
	fmt.Printf("Ukuran data dalam slice: %d byte\n", dataSize)
	fmt.Printf("Ukuran total slice (header + data yang digunakan): %d byte\n", totalSizeInBytes)
	fmt.Printf("Ukuran total kapasitas slice (header + kapasitas yang dialokasikan): %d byte\n", totalCapacitySizeInBytes)

	fmt.Println("\ndata2")
	fmt.Println(data2)
	// Menghitung ukuran slice dalam byte
	headerSize2 := unsafe.Sizeof(data2)
	elementSize2 := unsafe.Sizeof(data2[0])
	dataLength2 := len(data2)
	dataCapacity2 := cap(data2)
	dataSize2 := uintptr(dataLength) * elementSize2
	totalCapacitySize2 := uintptr(dataCapacity2) * elementSize2
	totalSizeInBytes2 := headerSize2 + dataSize2
	totalCapacitySizeInBytes2 := headerSize2 + totalCapacitySize2

	fmt.Printf("Ukuran header slice: %d byte\n", headerSize2)
	fmt.Printf("Ukuran elemen slice: %d byte\n", elementSize2)
	fmt.Printf("Panjang slice: %d\n", dataLength2)
	fmt.Printf("Kapasitas slice: %d\n", dataCapacity2)
	fmt.Printf("Ukuran data dalam slice: %d byte\n", dataSize2)
	fmt.Printf("Ukuran total slice (header + data yang digunakan): %d byte\n", totalSizeInBytes2)
	fmt.Printf("Ukuran total kapasitas slice (header + kapasitas yang dialokasikan): %d byte\n", totalCapacitySizeInBytes2)

}
