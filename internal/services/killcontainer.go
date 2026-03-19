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

    userid,getuseriderr := postgres.Getuserid(username)

	if getuseriderr != nil {
		return errors.ErrInternalserver
	}

	delerr := postgres.DeleteInstances(killinstancename,username,userid)

	if delerr != nil {
		return errors.ErrInstancenotfound
	}

	killerr := docker.Killcontainer(killinstancename,username)

	if killerr != nil {
		 return errors.ErrInternalserver
	}
	return nil 
}