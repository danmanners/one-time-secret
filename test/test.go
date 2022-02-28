package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func secret() ([]byte, error) {
	key := make([]byte, 32)

	if _, err := rand.Read(key); err != nil {
		return nil, err
	}

	return key, nil
}

func encrypt(plainData string, secret []byte) (string, error) {
	cipherBlock, err := aes.NewCipher(secret)
	if err != nil {
		return "", err
	}

	aead, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aead.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(aead.Seal(nonce, nonce, []byte(plainData), nil)), nil
}

func decrypt(encodedData string, secret []byte) (string, error) {
	encryptData, err := base64.URLEncoding.DecodeString(encodedData)
	if err != nil {
		return "", err
	}

	cipherBlock, err := aes.NewCipher(secret)
	if err != nil {
		return "", err
	}

	aead, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return "", err
	}

	nonceSize := aead.NonceSize()
	if len(encryptData) < nonceSize {
		return "", err
	}

	nonce, cipherText := encryptData[:nonceSize], encryptData[nonceSize:]
	plainData, err := aead.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plainData), nil
}

func main() {
	secret, err := secret()
	if err != nil {
		panic(fmt.Sprintf("unable to create secret key: %v", err))
	}

	message := "I am a very secret message"
	fmt.Println("Message:", message)

	encrypted, err := encrypt(message, secret)
	if err != nil {
		panic(fmt.Sprintf("unable to encrypt the data: %v", err))
	}
	fmt.Println("Encrypted:", encrypted)

	decrypted, err := decrypt(encrypted, secret)
	if err != nil {
		panic(fmt.Sprintf("unable to decrypt the data: %v", err))
	}
	fmt.Println("Decrypted:", decrypted)
}
