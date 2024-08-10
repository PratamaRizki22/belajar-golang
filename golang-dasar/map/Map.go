// map sifatnya key dan value, tidak ada index

package main

import "fmt"

func main() {
	employees := make(map[int]string)

	fmt.Println("empty map= ", employees)
	employees[12] = "joko"
	employees[15] = "susilo"

	for key, value := range employees {
		fmt.Println(key, ": ", value)
		fmt.Println(employees[key])
	}

}
