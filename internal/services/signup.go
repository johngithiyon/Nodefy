package services

import (
	"log"

	"github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/models"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/postgres"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/redis"
	"github.com/johngithiyon/Nodefy/pkg/utils"
)

func Signupservices(signup *models.Signup) (string,error) {

	     //check the existence 

		 checkexists :=  postgres.CheckUserexists(signup.Username)
		 if !checkexists  {
			   return "",errors.ErrAuthenticate
	   }

	    //valiadte username 

		 if !(len(signup.Username) <= 13)  {
			log.Println("Username Length Issues")
			return "",errors.ErrAuthenticate  
		 }

		 //validate email 

		 emailcheck := utils.Regexcompile()

		 if !emailcheck.MatchString(signup.Email) {
			log.Println("Email regex mismatch")
			return "",errors.ErrAuthenticate
		 }

		 //validate user password 
         
		 if !IsValidPassword(signup.Password) {
			    log.Println("Not a valid password")
			    return "",errors.ErrAuthenticate
		 }

		  //Hash the password 

		  hashedpass,hasherr := Passwordhash(signup.Password)

		  if hasherr != nil {
			 log.Println("Hash Err",hasherr)
			 return "",errors.ErrInternalserver
		  }

		  otp := OtpGenerator(signup.Email)

		if otp == "" {
			return "",errors.ErrInternalserver
		}

		senderr := SendEmail(signup.Email,otp)

		if senderr != nil {
			log.Println("Mail Send Error",senderr)
			return "",errors.ErrInternalserver
		}

		id := GenerateSessionStore(signup.Username)

		signupdetails := map[string]interface{} {
			"username":signup.Username,
			"email":signup.Email,
			"password":hashedpass,
			"otp":otp,
		}  

		storerr := redis.Storetemp(id,signupdetails)
		
        if  storerr != nil { 
			return "",errors.ErrInternalserver
	   } 

	   return id,nil 
}