package services

import (
	"encoding/base32"
	"log"
	"time"

	"github.com/pquerna/otp/totp"
)


func OtpGenerator(email string) string{

	  secretkey := email

	  bytekeys := []byte(secretkey)

	  encoded := base32.StdEncoding.EncodeToString(bytekeys)

       otp,otperr :=  totp.GenerateCode(encoded,time.Now())
	   
	   if otperr != nil {
		      log.Println(otperr)
			  return  ""
	   } 
	   return otp	   
}