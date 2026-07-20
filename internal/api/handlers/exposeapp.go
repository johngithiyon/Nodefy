package handlers

import (
	"net/http"

	"github.com/johngithiyon/Nodefy/internal/models"
	"github.com/johngithiyon/Nodefy/internal/services"
	"github.com/johngithiyon/Nodefy/pkg/response"
	"github.com/johngithiyon/Nodefy/pkg/utils"
)

func Exposeapphandler(w http.ResponseWriter,r *http.Request) {
              
	        var Deployform models.Exposeappform

			if r.Method != http.MethodPost {
				response.Response(w,405,"Invalid Metheod")
				return
		   }

		   sessionid,getsessionerr :=  utils.Getsessionid(r,"session-id")
   
		   if sessionid == "" && getsessionerr != nil {
				 response.Response(w,400,"Cookie not found")
				 return
			}
          
		   decoderr := utils.Jstost(r.Body,&Deployform)
   
		   if decoderr != nil {
			  response.Response(w,500,"Internal Server Error")
			  return 
		   }

		   exposerr := services.Exposeapp(sessionid,&Deployform)

		   if exposerr != nil {
			     response.Response(w,500,"Internal Server Error")
			     return
		   }

}