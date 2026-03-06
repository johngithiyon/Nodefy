package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/johngithiyon/Nodefy/internal/models"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/postgres"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/redis"
	"github.com/johngithiyon/Nodefy/internal/services"
	"github.com/johngithiyon/Nodefy/pkg/response"
	"github.com/johngithiyon/Nodefy/pkg/utils"
)

func Signuphandler(w http.ResponseWriter, r *http.Request) {

	      if r.Method != http.MethodPost {
			      log.Println("Invalid Metheod")
				  response.Response(w,405,"Invalid Metheod Request")
				  return 
		  }
	     
	      var Signup models.Signup
	           
	     Signup.Username =  r.FormValue("username")
		 Signup.Email = r.FormValue("email")
		 Signup.Password = r.FormValue("password")

		 //check the existence 

		    checkexists :=  postgres.CheckUserexists(&Signup)
			 if !checkexists  {
				  response.Response(w,409,"User Exists")
				  return 
			  }

		 //valiadte username 

		 if !(len(Signup.Username) <= 13)  {
			log.Println("Username Length Issues")
			response.Response(w,402,"Invalid Username")
			return 		  
		 }

		 //validate email 

		 emailcheck := utils.Regexcompile()

		 if !emailcheck.MatchString(Signup.Email) {
			log.Println("Email regex mismatch")
			response.Response(w,402,"Invalid Email")
			return 
		 }

		 //validate user password 
         
		 if !services.IsValidPassword(Signup.Password) {
			    log.Println("Not a valid password")
				response.Response(w,402,"Invalid Password")
			    return 
		 }

		 //Hash the password 

		 hashedpass,hasherr := services.Passwordhash(Signup.Password)

		 if hasherr != nil {
            log.Println("Hash Err",hasherr)
			response.Response(w,500,"Internal Server Error")
			return 
		 }


		otp := services.OtpGenerator(Signup.Email)

		if otp == "" {
			response.Response(w,500,"Internal Server Error")
			return 
		}

		senderr := services.SendEmail(Signup.Email,otp)

		if senderr != nil {
			log.Println("Mail Send Error",senderr)
			response.Response(w,500,"Internal Server Error")
			return 
		}

		id := services.GenerateSessionStore(Signup.Username)
		
		//store the temp id in cookies 

		http.SetCookie(w,&http.Cookie{

			Name: "temp-id",
			 Value: id,
			 Expires: time.Now().Add(30 * time.Minute),
	
	 })

		signupdetails := map[string]interface{} {
			"username":Signup.Username,
			"email":Signup.Email,
			"password":hashedpass,
			"otp":otp,
		}  

		storerr := redis.Storetemp(id,signupdetails)
		
        if  storerr != nil { 
            response.Response(w,500,"Internal Server Error")
			return
	   } 

		response.Response(w,200,"Signup Successful")

}