package main

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {
	password := "password123"
	hash256 := sha256.Sum256([]byte(password))
	hash512 := sha512.Sum512([]byte(password))

	fmt.Println(password)
	fmt.Println(hash256)
	fmt.Println(hash512)

	fmt.Printf("SHA-256 Hash hex value is %x\n", hash256)
	fmt.Printf("SHA-512 Hash hex value is %x\n", hash512)

	// -----------------------

	salt, err := generateSalt()
	fmt.Println("Original salt:", salt)
	fmt.Printf("Original salt hex: %x\n", salt)
	if err != nil {
		fmt.Println("Error generating salt:", err)
		return
	}

	// Hash the password with salt
	signupHash := hashPassword(password, salt)

	// Store the salt and password in database, just prini=ting as of now
	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Salt:", saltStr)
	fmt.Println("Hash:", signupHash)

	hashOriginalPassword := sha256.Sum256([]byte(password))
	fmt.Println("Hash of just the password string without salt:", base64.StdEncoding.EncodeToString(hashOriginalPassword[:]))

	// verify
	// retrieve the saltStr and dexode it
	decodedSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println("Unable to decode salt:", err)
		return
	}
	loginHash := hashPassword("password", decodedSalt)

	// compare the stored signupHash with loginHash

	if signupHash == loginHash {
		fmt.Println("Password is correct. you are loggedin.")
	} else {
		fmt.Println("Login failed. Please check your credentials.")
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

// func to hash password

func hashPassword(password string, salt []byte) string {
	saltedPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}
