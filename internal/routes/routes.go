package routes

import (
	"net/http"

	"github.com/johngithiyon/Nodefy/internal/api/handlers"
)

func Routes() {
	 
	  http.HandleFunc("/deploy",handlers.Deploy)
	  http.HandleFunc("/signup",handlers.Signuphandler)
}