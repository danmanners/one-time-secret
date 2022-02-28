package functions

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

// LOTS of logic was taken from this post:
// http://www.inanzzz.com/index.php/post/f3pe/data-encryption-and-decryption-with-a-secret-key-in-golang
func CheckForAesKey(Aes256Key string) {
	if Aes256Key == "rand" {
		randaeskey, err := generate_bytestring()
		if err != nil {
			panic(fmt.Sprintf("unable to create secret key: %v", err))
			os.Exit(1)
		}
		Aes256Key = string(randaeskey)
		fmt.Printf("Random AES256 Bytes Generated\n")
	} else {
		fmt.Printf("AES256 Bytes provided by environment variable\n")
	}
}

func Encrypt(plainData string, secret []byte) (string, error) {
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

func Decrypt(encodedData string, secret []byte) (string, error) {
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

func generate_bytestring() ([]byte, error) {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}
	return key, nil
}
