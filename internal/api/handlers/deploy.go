package handlers

import (
	"log"
	"net/http"

	"github.com/johngithiyon/Nodefy/internal/models"
	"github.com/johngithiyon/Nodefy/internal/services"

	"github.com/johngithiyon/Nodefy/pkg/utils"
)

func Deploy(w http.ResponseWriter, r *http.Request) {

	var Deploy models.Deploy

    if r.Method != http.MethodPost {
		 return
	}

	// Convert the Json into Struct

   decoderr := utils.Jsoncov(r.Body,&Deploy)

   if decoderr != nil {
	  log.Println(decoderr)
	  return 
   }

   // validate the user request

   if Deploy.OsName != "ubuntu" && Deploy.OsName != "alphine" {
	    
	       log.Println("Invalid os")
	       return 
   }

    
   deployerr := services.DeployInstances(&Deploy)

   if deployerr != nil {
	 return 
   }

}