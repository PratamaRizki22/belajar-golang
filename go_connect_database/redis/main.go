package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	// Membuat koneksi ke Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Alamat Redis
		Password: "",               // Kosong jika tidak ada password
		DB:       0,                // Gunakan database default (DB 0)
	})

	// Menyimpan beberapa key dan value ke Redis
	keyValues := map[string]string{
		"name":    "John Doe",
		"age":     "30",
		"country": "USA",
		"city":    "New York",
	}

	// Menyimpan setiap pasangan key-value
	for key, value := range keyValues {
		err := rdb.Set(ctx, key, value, 0).Err()
		if err != nil {
			log.Fatalf("Gagal menyimpan key %s: %v", key, err)
		}
		fmt.Printf("Key '%s' dengan Value '%s' disimpan di Redis\n", key, value)
	}

	// Mengambil dan menampilkan nilai dari Redis
	for key := range keyValues {
		val, err := rdb.Get(ctx, key).Result()
		if err != nil {
			if err == redis.Nil {
				fmt.Printf("Key '%s' tidak ditemukan di Redis\n", key)
			} else {
				log.Fatalf("Gagal mendapatkan key '%s': %v", key, err)
			}
		} else {
			fmt.Printf("Key '%s' memiliki Value '%s'\n", key, val)
		}
	}
}
