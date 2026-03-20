package services

import (
	"github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/repository/docker"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/postgres"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/redis"
	"github.com/johngithiyon/Nodefy/pkg/utils"
)

func Workspaceservices(sessionid string,id string) (string,error){

	instancename,getinstancerr := postgres.Getinstancebyid(id)

	if getinstancerr != nil {
		return "",getinstancerr
	}

	username,getusernamerr := redis.Getusername(sessionid)
 
	if getusernamerr != nil {
		return "",errors.ErrInternalserver
	}

	url,checkerr := postgres.CheckWorkspace(username)

	if checkerr != nil {
          return  "",errors.ErrInternalserver
	}

	if url == "" {
 
	// Docker call to start the workspace 

	urls,workspacerr := docker.Startworkspace(username,instancename)

	if workspacerr != nil {
		return "",errors.ErrInternalserver
	}

	url,converr := utils.Jsonconvertor(urls)

	if converr != nil {
		return "",errors.ErrInternalserver
	}

	saverr := postgres.SaveWorkspace(username,url)

	if saverr != nil {
		 return "",errors.ErrInternalserver
	}

	return url,nil 

}  

 return  url,nil 

  
}