package main

import (
	"log"
	"net/http"

	"github.com/johngithiyon/Nodefy/internal/app"

)

func main() {
	 
	 app.Appstartup()
	 log.Println("Server Is  listening") 
	 http.ListenAndServe(":8080",nil)
}