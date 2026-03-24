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

func Buildhandler(w http.ResponseWriter, r *http.Request) {

	var Build  models.Build

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

   decoderr := utils.Jstost(r.Body,&Build)

   if decoderr != nil {
	  log.Println("Decode Error",decoderr)
	  response.Response(w,500,"Internal Server Error")
	  return 
   }

    
   builderr := services.BuildInstances(sessionid,&Build)

   if builderr == errors.ErrInternalserver {
	    response.Response(w,500,"Internal Server Error")
		return
   }

   if builderr == errors.ErrBadrequest {
	     response.Response(w,400,"Bad Request")
		 return
   }

   response.Response(w,200,"Instances Build Successfully")

}