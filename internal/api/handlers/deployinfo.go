package handlers

import (
	"net/http"

	"github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/services"
	"github.com/johngithiyon/Nodefy/pkg/response"
	"github.com/johngithiyon/Nodefy/pkg/utils"
)

func GetDeployinstance(w http.ResponseWriter, r *http.Request) {
	 
	if r.Method != http.MethodGet {
		response.Response(w,405,"Invalid Method Type")
		return 
   }
  
  
   sessionid,sessionerr := utils.Getsessionid(r,"session-id")

   if sessionerr != nil {
		response.Response(w,400,errors.ErrBadrequest.Error())
		return 
   }

  resp,getinfoerr :=  services.Getdeployinfo(sessionid)


	if getinfoerr != nil {
		response.Response(w,500,"Internal Server Error")
		return 
	}

	w.Write([]byte(resp))

    
}