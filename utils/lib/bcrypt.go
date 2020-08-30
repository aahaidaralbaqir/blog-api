package lib

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func Hash(password string) string{
	var toByte []byte = []byte(password)
	hash,err := bcrypt.GenerateFromPassword(toByte,bcrypt.DefaultCost)

	if err != nil {
		log.Println(err)
	}

	return string(hash)

}

func CheckHash(hashed string, plain string) bool {
	hashedToByte := []byte(hashed)
	plaintToByte := []byte(plain)
	err := bcrypt.CompareHashAndPassword(hashedToByte,plaintToByte)

	if err != nil {
		return false
	}
	return true
}