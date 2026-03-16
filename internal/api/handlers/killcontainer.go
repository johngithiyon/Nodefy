package handlers

import (
	"net/http"

	"github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/services"
	"github.com/johngithiyon/Nodefy/pkg/response"
	"github.com/johngithiyon/Nodefy/pkg/utils"
)

func Killcontainerhandler(w http.ResponseWriter, r *http.Request) {

	    killinstancename := r.FormValue("killinstancename")     
	
		if r.Method != http.MethodPost {
			response.Response(w,405,"Invalid Metheod")
			return 
		}

	    //get the username to stop the container

		sessionid,getsessionerr :=  utils.Getsessionid(r,"session-id")

		if sessionid == "" && getsessionerr != nil {
			 response.Response(w,400,"Cookie not found")
			 return
		}


        killerr := services.Killcontainerservices(killinstancename,sessionid)

		if killerr == errors.ErrInternalserver {
			response.Response(w,500,"Internal Server Error")
			return 
		}

		if killerr == errors.ErrInstancenotfound {
			response.Response(w,400,errors.ErrInstancenotfound.Error())
			return
		}

		response.Response(w,200,"Kill Instance Successfully")

}