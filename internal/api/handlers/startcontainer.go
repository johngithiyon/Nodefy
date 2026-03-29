package handlers

import (
	"net/http"

	"github.com/johngithiyon/Nodefy/internal/models"
	"github.com/johngithiyon/Nodefy/internal/services"
	"github.com/johngithiyon/Nodefy/pkg/response"
	"github.com/johngithiyon/Nodefy/pkg/utils"
)

func Startcontainerhandler(w http.ResponseWriter, r *http.Request) {

	     var startcontainer models.Containermanage

		 if r.Method != http.MethodPost {
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
   
	  decoderr := utils.Jstost(r.Body,&startcontainer)
   
	  if decoderr != nil {
		 response.Response(w,500,"Internal Server Error")
		 return 
	  }

	 starterr :=  services.Startcontainerserives(sessionid,&startcontainer)

	 if starterr != nil {
		response.Response(w,500,"Internal Server Error")
		return 
	 }

	 response.Response(w,200,"Start Container Sucessfully")
}