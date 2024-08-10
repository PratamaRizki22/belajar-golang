package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Going to sleep for 5 seconds")
	time.Sleep(5 * time.Second)
	fmt.Println("I woke up after 5 seconds")
}
