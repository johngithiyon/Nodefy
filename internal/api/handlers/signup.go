package handlers

import (
	"log"
	"net/http"

	"github.com/johngithiyon/Nodefy/internal/models"
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

		 //valiadte username 

		 if !(len(Signup.Username) <= 12)  {
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

		 hashedpass,hasherr := services.Passwordhash(Signup.Password)

		 if hasherr != nil {
			return 
		 }

		 log.Println(hashedpass)

		 response.Response(w,200,"Signup Successful")

         


}