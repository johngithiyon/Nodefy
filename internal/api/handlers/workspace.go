package handlers

import (
	"net/http"

	"github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/services"
	"github.com/johngithiyon/Nodefy/pkg/response"
	"github.com/johngithiyon/Nodefy/pkg/utils"
)

func WorkspaceHanlder(w http.ResponseWriter, r *http.Request) {


	        id := r.URL.Query().Get("id")

			sessionid,sessionerr := utils.Getsessionid(r,"session-id")

			if sessionerr != nil {
				 response.Response(w,400,errors.ErrBadrequest.Error())
				 return 
			}

			url,workspacerr := services.Workspaceservices(sessionid,id)

			if workspacerr != nil {
				response.Response(w,500,"Internal Server Error")
				return
			}

			w.Write([]byte(url))
}