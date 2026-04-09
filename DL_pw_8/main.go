package main

import (
	rsa "DL_pw_8/rsa"
	"fmt"
	"math/big"
	"strconv"
)

func main() {

	var input string

	fmt.Println("Enter a message to encrypt(integer only):")
	fmt.Scanln(&input)

	if _, err := strconv.Atoi(input); err != nil {
		fmt.Println("The message is not an integer.")
	} else {
		message, _ := new(big.Int).SetString(input, 10)

		publicKey, privateKye := rsa.KeyGen()

		encrypted_message := rsa.Encrypt(message, &publicKey)
		decrypted := rsa.Decrypt(encrypted_message, &privateKye)

		fmt.Printf("Input data: %d, encrypted data: %s, decrypted data: %s\n", message, encrypted_message, decrypted)
	}

}
