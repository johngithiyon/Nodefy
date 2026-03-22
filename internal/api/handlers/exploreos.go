package handlers

import (
	"net/http"

	"github.com/johngithiyon/Nodefy/internal/models"
	"github.com/johngithiyon/Nodefy/pkg/response"
	"github.com/johngithiyon/Nodefy/pkg/utils"
)

func ExploreosHandler(w http.ResponseWriter, r *http.Request) {
	 
	       var exploreos models.Exploreos

		   if r.Method != http.MethodPost {
			response.Response(w,405,"Invalid Metheod")
			return
	   }


	   sessionid,getsessionerr :=  utils.Getsessionid(r,"session-id")

	   if sessionid == "" && getsessionerr != nil {
		     response.Response(w,400,"Cookie not found")
			 return
		}

        
	// Convert the Json into Struct

   decoderr := utils.Jstost(r.Body,&exploreos)

   if decoderr != nil {
	  response.Response(w,500,"Internal Server Error")
	  return 
   }



}