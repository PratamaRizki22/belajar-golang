package main

import (
	"fmt"
	"net/http"
)

func contentType(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	cType := res.Header.Get("Content-Type")

	return cType, nil

}

func main() {
	url := "http://self-learning-java-tutorial.blogspot.com"

	contrntType, err := contentType(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Content Type: ", contrntType)
}
