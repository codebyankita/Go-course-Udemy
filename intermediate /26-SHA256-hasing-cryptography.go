package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {

	password := "password123"

	// hash := sha256.Sum256([]byte(password))
	// hash512 := sha512.Sum512([]byte(password))

	// fmt.Println(password)
	// fmt.Println(hash)
	// fmt.Println(hash512)
	// fmt.Printf("SHA-256 Hash hex val: %x\n", hash)
	// fmt.Printf("SHA-512 Hash hex val: %x\n", hash512)

	salt, err := generateSalt()
	fmt.Println("Original Salt:", salt)
	fmt.Printf("Original Salt: %x\n", salt)
	if err != nil {
		fmt.Println("Error generating salt:", err)
		return
	}

	// Hash the password with salt
	signUpHash := hashPassword(password, salt)

	// Store the salt and password in database, just printing as of now
	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Salt:", saltStr)    // simulate as storing in database
	fmt.Println("Hash:", signUpHash) // simulate as storing in database
	hashOriginalPassword := sha256.Sum256([]byte(password))
	fmt.Println("Hash of just the password string without salt:", base64.StdEncoding.EncodeToString(hashOriginalPassword[:]))

	// verify
	// retrieve the saltStr and decode it
	decodedSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println("Unable to decode salt:", err)
		return
	}
	loginHash := hashPassword("password124", decodedSalt)

	// Compare the stored signUpHash with loginHash
	if signUpHash == loginHash {
		fmt.Println("Password is correct. You are logged in.")
	} else {
		fmt.Println("Login failed. Please check user credentials.")
	}

}

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

// Function to hash password
func hashPassword(password string, salt []byte) string {
	saltedPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}
