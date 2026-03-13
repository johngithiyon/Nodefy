package routes

import (
	"net/http"

	"github.com/johngithiyon/Nodefy/internal/api/handlers"
)

func Routes() {
	 
	  http.HandleFunc("/deployinstance",handlers.Deploy)
	  http.HandleFunc("/signup",handlers.Signuphandler)
	  http.HandleFunc("/otp",handlers.Otphandler)
	  http.HandleFunc("/login",handlers.Loginhandler)
	  http.HandleFunc("/stopinstance",handlers.StopcontainerHandler)
	  http.HandleFunc("/killinstance",handlers.Killcontainerhandler)
}