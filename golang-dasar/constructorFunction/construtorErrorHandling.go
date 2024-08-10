package main

import (
	"errors"
	"fmt"
)

type Ciri struct {
	namaHewan  string
	tipe       string
	jumlahKaki int8
	makanan    [3]string
}

func (animal *Ciri) createAnimal(namaHewan string, tipe string, jumlahKaki int8, contoh [3]string) error {
	// animal := &Ciri{}

	if jumlahKaki < 0 {
		return errors.New("jumlah kaki harus lebih dari 0")
	}
	animal.namaHewan = namaHewan
	animal.tipe = tipe
	animal.jumlahKaki = jumlahKaki
	animal.makanan = contoh
	return nil
}

func main() {
	ayam := &Ciri{}
	error1 := ayam.createAnimal("Ayam", "Mamalia", -1, [3]string{"cacing", "beras", "pur"})

	if error1 != nil {
		fmt.Println("error: ", error1)
		return
	}

	fmt.Printf(`
nama hewan: %v
tipee: %v
jumlah kaki: %d
makanan: %v
	`, ayam.namaHewan, ayam.tipe, ayam.jumlahKaki, ayam.makanan)
}
