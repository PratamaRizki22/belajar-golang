package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func encrypt(plaintext, key []byte) (ciphertext, nonce []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}

	nonce = make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, nil, err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}

	ciphertext = aesGCM.Seal(nil, nonce, plaintext, nil)
	return ciphertext, nonce, nil

}
