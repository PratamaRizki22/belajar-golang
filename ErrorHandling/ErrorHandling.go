package main

import "fmt"

func main() {
	result, err1 := add(10, 20)

	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println("sum 10 and 20: ", result)
	}

	result2, err2 := add(0, -1)

	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("sum of 0 and -1: ", result2)
	}
}

func add(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, fmt.Errorf("sum pf negative number is not supported")
	}

	return a + b, nil
}
