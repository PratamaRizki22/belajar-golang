// sumber kode: https://www.youtube.com/watch?v=zPYjfgxYO7k&t=28s

// sumber api: https://github.com/emsifa/api-wilayah-indonesia
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

type Region struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var regions []Region

func main() {
	res, err := http.Get("https://www.emsifa.com/api-wilayah-indonesia/api/provinces.json")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("response tidak tersedia")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &regions)
	if err != nil {
		panic(err)
	}

	for _, region := range regions {
		fmt.Println(region.ID, region.Name)
	}

	// tebak id
	rand.Shuffle(len(regions), func(i, j int) { regions[i], regions[j] = regions[j], regions[i] })

	fmt.Println("tebak nama provinsi berdasarkan ID-nya!")
	fmt.Println("Ketik 'exit' untuk keluar dari game.")

	reader := bufio.NewReader(os.Stdin)
	for _, region := range regions {
		// options := generateOption
		fmt.Printf("apa nama provinsi dengan id: %s? negara: %s\n", region.ID, region.Name)

		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)

		if strings.ToLower(answer) == "exit" {
			fmt.Println("Terima kasih sudah bermain!")
			break
		}

		if strings.EqualFold(answer, region.Name) {
			fmt.Println(" anda benar gg")

		} else {
			fmt.Printf("salah coooyy yang benar %s\n", region.Name)
		}
	}
}

// func generateOption(correct)
