package handlers

import (
	"log"
	"net/http"

	"github.com/johngithiyon/Nodefy/internal/models"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/redis"
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


	username,geterr := redis.Getusername(sessionid)

	if geterr != nil {
		response.Response(w,500,"Internal Server Error")
		return
	}

	// Convert the Json into Struct

   decoderr := utils.Jsoncov(r.Body,&Deploy)

   if decoderr != nil {
	  log.Println("Decode Error",decoderr)
	  response.Response(w,500,"Internal Server Error")
	  return 
   }

   // validate the user request

   if Deploy.OsName != "ubuntu" && Deploy.OsName != "alphine" {
	    
	       log.Println("Invalid os")
		   response.Response(w,400,"Invalid OS")
	       return 
   }

   // validate the user selected services

   for _,services := range Deploy.Services {
	     
	       if services != "postgresql" && services != "redis-server" {

				log.Println("Invalid Services")
				response.Response(w,400,"Invalid Services")
				return 
		   }
   }

    
   deployerr := services.DeployInstances(username,&Deploy)

   if deployerr != nil {
	  log.Println("Deploy Error",deployerr)
	  response.Response(w,500,"Internal Server Error")
	  return 
   }

   response.Response(w,200,"Instances Deployed Successfully")

}