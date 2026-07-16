package main

import (
	"log"
	"net/http"

	"github.com/johngithiyon/Nodefy/internal/app"

)

func main() {
	 
	 app.Appstartup()
	 log.Println("Server Is  listening") 
	 listenerr := http.ListenAndServe("[2409:40f4:1000:6af0:a582:89ed:b742:ee12]:8080",nil)

	 if listenerr != nil {
		 log.Println("Listen err",listenerr)
	 }
}