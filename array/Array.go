package main

//  aray tidak bisa diubah atau immutable
import "fmt"

func UpdateMakanan(makananbaru [3]string) [3]string {
	makananbaru[0] = "bakso"
	makananbaru[1] = "mie ayam"
	makananbaru[2] = "geprek"

	return makananbaru
}

func main() {
	var countrie [4]int
	countries := [5]string{}

	countrie[3] = 2
	countries[0] = "jawa"
	countries[1] = "indoneisa"

	fmt.Println(countries, countrie)
	fmt.Println("len countrie: ", len(countries))

	indonesia := [4]int8{1, 2, 3, 4}
	fmt.Println(indonesia[0])

	for i := 0; i < len(indonesia); i++ {
		fmt.Println(indonesia[i])
	}

	for i := 0; i < len(countries[0]); i++ {
		fmt.Println(string(countries[0][i]))
	}

	fmt.Print("pemmbuktian immutable\n")

	makanan := [3]string{"soto", "nasi goreng", "ayam"}

	beliBaru := UpdateMakanan(makanan)

	// tidak bisa diubah, tapi bisa buat variable baru
	fmt.Println("lama: ", makanan)
	fmt.Println("variable baru: ", beliBaru)

	makan, ok := indonesia[0], indonesia[2]
	fmt.Println(makan)
	fmt.Println(ok)

}
