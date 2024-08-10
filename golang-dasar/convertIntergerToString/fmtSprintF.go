package main

import (
	"fmt"
	"reflect"
)

func main() {
	str := fmt.Sprintf("%v", 123)
	fmt.Println(str, reflect.TypeOf(str))
}
