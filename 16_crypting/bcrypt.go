package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)


func main() {
	pass := "example_password"

	hash := hashPassword(pass)
	fmt.Println("[]byte", []byte(hash))
	fmt.Println("string", string(hash))

	fmt.Println(comparePasswordAndHash(pass, []byte(hash)))
}

func hashPassword(pass string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 11)
	if (err == nil) {
		return string(hash)
	} else {
		return "err"
	}
}

func comparePasswordAndHash(pass string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(pass))
	return err == nil
}