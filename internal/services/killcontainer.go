package services

import (
	"log"

	"github.com/johngithiyon/Nodefy/internal/repository/docker"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/redis"
)

func Killcontainerservices(sessionid string) error{

	username,getusernamerr := redis.Getusername(sessionid)

	if username == "" && getusernamerr != nil {
		  log.Println("Username get err in kill container",getusernamerr)
		  return getusernamerr
	}

	  killerr :=  docker.Killcontainer(username)

	  if killerr != nil {
		  return killerr
	  }
     
	   return nil 
}