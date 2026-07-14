package services

import (

	"github.com/johngithiyon/Nodefy/internal/repository/docker"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/postgres"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/redis"
	"github.com/johngithiyon/Nodefy/pkg/utils"
)

func Createworkspace(sessionid string) (string,error) {

	       username,geterr := redis.Getusername(sessionid)

		   if geterr != nil {
			   return "",geterr
		   }

		   containername,workspacerr := docker.Startworkspace(username)

			if workspacerr != nil {
				return "",workspacerr
			}
	  
			id,generaterr := utils.Generateworkspaceid()

			if generaterr != nil {
				  return "",generaterr
			}
		    
	        saverr := postgres.SaveWorkspace(username,id,containername)

			if saverr != nil {
				return "",saverr
			}

			return id,nil 
}