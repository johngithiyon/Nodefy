package handlers

import (
	"log"
	"net/http"

	"github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/models"
	"github.com/johngithiyon/Nodefy/internal/services"

	"github.com/johngithiyon/Nodefy/pkg/response"
	"github.com/johngithiyon/Nodefy/pkg/utils"
)

func Deploy(w http.ResponseWriter, r *http.Request) {

	var Deploy models.Deploy

    if r.Method != http.MethodPost {
		 log.Println("Invalid Metheod")
		 response.Response(w,405,"Invalid Metheod")
		 return
	}

	//get the username to find the deploy user 
 
	sessionid,getsessionerr :=  utils.Getsessionid(r,"session-id")

	   if sessionid == "" && getsessionerr != nil {
		     response.Response(w,400,"Cookie not found")
			 return
		}


	// Convert the Json into Struct

   decoderr := utils.Jsoncov(r.Body,&Deploy)

   if decoderr != nil {
	  log.Println("Decode Error",decoderr)
	  response.Response(w,500,"Internal Server Error")
	  return 
   }

    
   deployerr := services.DeployInstances(sessionid,&Deploy)

   if deployerr == errors.ErrInternalserver {
	    response.Response(w,500,"Internal Server Error")
		return
   }

   if deployerr == errors.ErrBadrequest {
	     response.Response(w,400,"Bad Request")
		 return
   }

   response.Response(w,200,"Instances Deployed Successfully")

}