package services

import (
	"github.com/johngithiyon/Nodefy/internal/repository/storage/postgres"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/redis"
)

func Getworkspace(sessionid string) (string,error) {
 
	  username,geterr := redis.Getusername(sessionid)

	  if geterr != nil {
		    return "",geterr
	  }

	  workspaceurl,workspaceurlgeterr := postgres.Getworkspace(username)

	  if workspaceurlgeterr != nil {
		    return "",workspaceurlgeterr
	  }
	
	   return workspaceurl,nil 
}