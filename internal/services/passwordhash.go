package services

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Passwordhash(password string ) (string,error){

	// Use bcrypt for the password hash 
	  
	password_hash,hasherr := bcrypt.GenerateFromPassword([]byte(password),10)

	if hasherr != nil {
		  //resp.JsonError(w,"Internal Server Error")
		  log.Println("Cannot generate the hash value",hasherr)
		  return "",hasherr 
 
	}	
	
	return string(password_hash),nil 
}