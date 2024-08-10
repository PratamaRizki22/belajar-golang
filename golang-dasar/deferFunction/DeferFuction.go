package main

//defer dieksekusi setelah semua program dijalankan
import "fmt"

func main() {
	defer fmt.Print(", Wellcome golanv\n")
	defer fmt.Print("Hello")
	fmt.Println("saya, koko")
}
