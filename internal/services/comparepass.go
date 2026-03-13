package services

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Comparepassword(userpass string,passwd string) error {
          
	err := bcrypt.CompareHashAndPassword([]byte(passwd),[]byte(userpass))
	
	if err != nil {
		log.Println("Password Mismatch")
		return err
	}

	return nil 
}