package handlers

import (
	"log"
	"net/http"

	"github.com/johngithiyon/Nodefy/internal/services"
	"github.com/johngithiyon/Nodefy/pkg/response"
	"github.com/johngithiyon/Nodefy/pkg/utils"
)

func Createworkspace(w http.ResponseWriter,r *http.Request) {
	          
	if r.Method != http.MethodGet {
		log.Println("Invalid Metheod")
		response.Response(w,405,"Invalid Metheod")
		return
   }  

   sessionid,getsessionerr :=  utils.Getsessionid(r,"session-id")
   
   if sessionid == "" && getsessionerr != nil {
		 response.Response(w,400,"Cookie not found")
		 return
	}

   id,createrr := services.Createworkspace(sessionid)

   if createrr != nil {
		response.Response(w,500,"Internal Server Error")
		return 
   }

   url := "http://"+id+".workspace.nodefy.in"

   response.Response(w,200,url)
}