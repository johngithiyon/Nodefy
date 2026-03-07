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
          
		  passwd,passwderr := postgres.SearchPassword(login.Username)

          if  passwd == "" && passwderr != nil {
                response.Response(w,401,"Login Failed")
				return 
		  }

		 comparerr :=  services.Comparepassword(login.Password,passwd)

		 if comparerr != nil {
			  response.Response(w,401,"Username or Password Mismatch")
			  return 
		 }

		 id := services.GenerateSessionStore(login.Username)

		 http.SetCookie(w,&http.Cookie{

			 Name: "session-id",
			 Value: id,
			 Expires: time.Now().Add(7 * 24 * time.Hour),
	 })

	     seterr :=  redis.Storesessionid(id,login.Username)

		 if seterr != nil {
			  response.Response(w,500,"Internal Server Error")
			  return 
		 }
		 
		  response.Response(w,200,"Login Successfull")
}