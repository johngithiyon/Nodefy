package handlers

import (
	"net/http"

	"github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/services"
	"github.com/johngithiyon/Nodefy/pkg/response"
	"github.com/johngithiyon/Nodefy/pkg/utils"
)

func GetInstancesHandler(w http.ResponseWriter,r *http.Request) {

	     if r.Method != http.MethodGet {
			  response.Response(w,405,"Invalid Method Type")
			  return 
		 }
        
	    
	     sessionid,sessionerr := utils.Getsessionid(r,"session-id")

		 if sessionerr != nil {
			  response.Response(w,400,errors.ErrBadrequest.Error())
              return 
         }

         instances,getinstancerr := services.Getinstanceservices(sessionid)

		 if getinstancerr != nil {
			return
		 }

		w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(instances))
}		 
