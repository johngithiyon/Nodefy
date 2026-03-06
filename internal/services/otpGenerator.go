package services

import (
	"encoding/base32"
	"log"
	"time"

	"github.com/johngithiyon/Nodefy/internal/repository/storage/redis"
	"github.com/pquerna/otp/totp"
)


func OtpGenerator(email string) string{

	  secretkey := email

	  bytekeys := []byte(secretkey)

	  encoded := base32.StdEncoding.EncodeToString(bytekeys)

       otp,otperr :=  totp.GenerateCode(encoded,time.Now())
	   
	   if otperr != nil {
		      log.Println(otperr)
	   } 

	  seterr := redis.SetOtp(email,otp)

	  if seterr != nil {
		   log.Println("Seterr otp",seterr)
		   return ""
	  }

	   return otp	   
}