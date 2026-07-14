package handlers

import (
	"net/http"

	"github.com/johngithiyon/Nodefy/pkg/utils"
)

func Roothandler(w http.ResponseWriter,r *http.Request) {
	      
	      host := r.Host

		  check := "workspace.nodefy.in"

		  checker := utils.Istringcheck(host,check)

		  if checker {
			      
			    WorkspaceHanlder(w,r)
			    
		  } else {
		   
		       Renderhomepage(w,r)

		  } 
		  
}