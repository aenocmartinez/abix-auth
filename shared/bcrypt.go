package shared

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println("shared / HashAndSalt / bcrypt.GenerateFromPassword: " + err.Error())
	}
	return string(hash)
}

func ComparePassword(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		fmt.Println(err)
		log.Println("shared / ComparePassword / bcrypt.CompareHashAndPassword: " + err.Error())
		return false
	}
	return true
}
