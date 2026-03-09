package services

import (
	"log"

	"github.com/johngithiyon/Nodefy/internal/repository/docker"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/redis"
)

func Stopcontainerservices(sessionid string) error{

	username,getusernamerr := redis.Getusername(sessionid)

	if username == "" && getusernamerr != nil {
		  log.Println("Get username err from the stop container",getusernamerr)
		  return getusernamerr
	}

	  dockerstoperr :=   docker.Stopcontainer(username)

	  if dockerstoperr != nil {
		    return dockerstoperr
	  }

	  return  nil 
}