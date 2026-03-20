package services

import (
	"html/template"
	"log"
	"net/http"
)


func Render(w http.ResponseWriter, name string) {

	      tmpl,tmplerr := template.ParseFiles("templates/"+name)

		  if tmplerr != nil {
			   log.Println("Cannot Parse the files",tmplerr)
			   return
		  }

		  tmpl.Execute(w,nil)
		  
}