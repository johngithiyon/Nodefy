package handlers

import (
	"net/http"

	"github.com/johngithiyon/Nodefy/internal/repository/storage/redis"
	"github.com/johngithiyon/Nodefy/internal/services"
	"github.com/johngithiyon/Nodefy/pkg/response"
	"github.com/johngithiyon/Nodefy/pkg/utils"
)

func StopcontainerHandler(w http.ResponseWriter, r *http.Request) {

	         if r.Method != http.MethodGet {
				 response.Response(w,405,"Invalid Metheod")
				 return 
			 }

	         //get the username to stop the container

			sessionid,getsessionerr :=  utils.Getsessionid(r,"session-id")

			if sessionid == "" && getsessionerr != nil {
				 response.Response(w,400,"Cookie not found")
				 return
			}

			username,getusernamerr := redis.Getusername(sessionid)

			if username == "" && getusernamerr != nil {
				  response.Response(w,500,"Internal Server Error")
				  return
			}

           dockerstoperr :=  services.Stopcontainerservices(username)

		   if dockerstoperr != nil {
			   response.Response(w,500,"Internal Server Error")
			   return
		   }

           response.Response(w,200,"Stop Instances Successfully")

}