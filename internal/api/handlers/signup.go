package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/models"
	"github.com/johngithiyon/Nodefy/internal/services"
	"github.com/johngithiyon/Nodefy/pkg/response"
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

		id,signuperr := services.Signupservices(&Signup)

		if signuperr == errors.ErrAuthenticate {
			response.Response(w,402,"Login Failed")
			return 
		}

		if signuperr == errors.ErrInternalserver {
			response.Response(w,500,"Internal Server Error")
			return
		}
		
		//store the temp id in cookies 

		http.SetCookie(w,&http.Cookie{

			Name: "temp-id",
			 Value: id,
			 Expires: time.Now().Add(30 * time.Minute),
	
	 })

		response.Response(w,200,"Signup Successful")

}