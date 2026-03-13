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

func Loginhandler(w http.ResponseWriter, r *http.Request) {
	 
	      if r.Method != http.MethodPost {
			log.Println("Invalid Metheod")
			response.Response(w,405,"Invalid Metheod Request")
			return 
		  }

		  var login models.Login

		  login.Username = r.FormValue("username")
		  login.Password = r.FormValue("password")


         id,loginerr := services.LoginServices(login.Username,login.Password)
		 
		 if loginerr == errors.ErrAuthenticate {
			  response.Response(w,402,"Login Failed")
			  return 
		 }

		 if loginerr == errors.ErrInternalserver {
			 response.Response(w,500,"Internal Server Error")
			 return 
		 }
		 

		 http.SetCookie(w,&http.Cookie{

			 Name: "session-id",
			 Value: id,
			 Expires: time.Now().Add(7 * 24 * time.Hour),
	 })

		 
		  response.Response(w,200,"Login Successfull")
}