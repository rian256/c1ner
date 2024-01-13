package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func Encryptor(inputFile string, outputFile string, key []byte) error {
	fileContent, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	stream := cipher.NewCTR(block, iv)
	ciphertext := make([]byte, len(fileContent))
	stream.XORKeyStream(ciphertext, fileContent)

	outputFileContent := append(iv, ciphertext...)
	err = ioutil.WriteFile(outputFile, outputFileContent, 0644)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	key := []byte("Lorem Ipsum 5432")

	filePattern := "safas.txt"
	outputFile := "lohgabahayata.wleee"

	err := Encryptor(filePattern, outputFile, key)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Remove(filePattern)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File '%s' encrypted and renamed to '%s'\n", filePattern, outputFile)
}
