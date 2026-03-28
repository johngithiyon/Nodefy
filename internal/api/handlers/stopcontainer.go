package handlers

import (
	"net/http"

	"github.com/johngithiyon/Nodefy/internal/models"
	"github.com/johngithiyon/Nodefy/internal/services"
	"github.com/johngithiyon/Nodefy/pkg/response"
	"github.com/johngithiyon/Nodefy/pkg/utils"
)

func StopcontainerHandler(w http.ResponseWriter, r *http.Request) {

	         var Stopcontainer models.Stopcontainer

	         if r.Method != http.MethodGet {
				 response.Response(w,405,"Invalid Metheod")
				 return 
			 }

	         //get the username to stop the container

			sessionid,getsessionerr :=  utils.Getsessionid(r,"session-id")

			if sessionid == "" && getsessionerr != nil {
				 response.Response(w,400,"Cookie not found")
				 return
			}


			// Convert the Json into Struct

			decoderr := utils.Jstost(r.Body,&Stopcontainer)

			if decoderr != nil {
			response.Response(w,500,"Internal Server Error")
			return 
			}


           dockerstoperr :=  services.Stopcontainerservices(sessionid,&Stopcontainer)

		   if dockerstoperr != nil {
			   response.Response(w,500,"Internal Server Error")
			   return
		   }

           response.Response(w,200,"Stop Instances Successfully")

}