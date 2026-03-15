package services

import (
	"log"

	"github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/repository/docker"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/postgres"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/redis"
)

func Killcontainerservices(killinstancename string,sessionid string) error{

	username,getusernamerr := redis.Getusername(sessionid)

	if username == "" && getusernamerr != nil {
		  log.Println("Username get err in kill container",getusernamerr)
		  return errors.ErrInternalserver
	}

	res :=  postgres.CheckInstance(killinstancename,username)

	if res == 1 {

	  killerr :=  docker.Killcontainer(killinstancename,username)

	  if killerr != nil {
		  return killerr
	  }

	 rmerr :=  postgres.RemoveInstance(killinstancename,username)

	 if rmerr != nil {
		 return  rmerr
	 }
 
	 return nil 

	}  else {
		
		return  errors.ErrInstancenotfound
	}
}