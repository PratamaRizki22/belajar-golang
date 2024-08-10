package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	str := strconv.FormatInt(int64(123), 10)
	fmt.Println(str, reflect.TypeOf(str))
}
