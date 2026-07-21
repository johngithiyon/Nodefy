package handlers

import (
	"log"
	"net/http"

	"github.com/johngithiyon/Nodefy/internal/services"
	"github.com/johngithiyon/Nodefy/pkg/response"
	"github.com/johngithiyon/Nodefy/pkg/utils"
)

func Getworkspace(w http.ResponseWriter, r *http.Request) {

	log.Println("get workspace")

	if r.Method != http.MethodGet {
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

	workspace_url,geterr := services.Getworkspace(sessionid)

	if geterr != nil {
		response.Response(w,500,"Internal Server Error")
		return
	}

	w.Write([]byte(workspace_url))
}