package handlers

import (
	"log"
	"net/http"

	"github.com/johngithiyon/Nodefy/internal/models"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/postgres"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/redis"
	"github.com/johngithiyon/Nodefy/pkg/response"
)


func Otphandler(w http.ResponseWriter, r *http.Request) {

	   var Ver models.Verify

	    if r.Method != http.MethodPost  {
			log.Println("Invalid Request Metheod")
			response.Response(w,405,"Invalid Request Metheod")
		}
	 
	    otp := r.FormValue("otp")

		id,cookierr :=  r.Cookie("temp-id")

		if cookierr != nil {
			log.Println("Cookies Error",cookierr)
			return 
		}

        storedata,geterr :=  redis.Getstoretemp(id.Value)

		if storedata == nil  && geterr != nil {
			 response.Response(w,500,"Internal Server Error")
			 return 
		}

		if otp == storedata["otp"] {

			Ver.Username = storedata["username"]
			Ver.Email = storedata["email"]
			Ver.Password = storedata["password"]

			storerr := postgres.StoreUser(&Ver)

			if storerr != nil {
				response.Response(w,500,"Internal Server Error")
				return 
			}
			response.Response(w,200,"Otp Verified")
			return 
		} else {
			 response.Response(w,401,"Otp Verify failed")
			 return
		}



}