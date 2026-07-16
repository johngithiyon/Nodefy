package handlers

import (
	"net/http"

	"github.com/johngithiyon/Nodefy/pkg/utils"
)

func Roothandler(w http.ResponseWriter,r *http.Request) {
	      
	      host := r.Host

		  //here give your subdomain name for the workspace

		  check := "workspace.nodefy.in:8080"

		  checker := utils.Istringcheck(host,check)

		  if checker {
			      
			    WorkspaceHanlder(w,r)
			    
		  } else {
		   
		       Renderhomepage(w,r)

		  } 
		  
}