package handlers

import (
	"log"
	"net/http"

	"github.com/johngithiyon/Nodefy/internal/models"
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

		  

		  response.Response(w,200,"Login Successfull")
}