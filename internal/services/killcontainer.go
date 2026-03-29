package services

import (
	"log"

	"github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/models"
	"github.com/johngithiyon/Nodefy/internal/repository/docker"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/postgres"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/redis"
)

func Killcontainerservices(killinstance *models.Containermanage,sessionid string) error{

	username,getusernamerr := redis.Getusername(sessionid)

	if username == "" && getusernamerr != nil {
		  log.Println("Username get err in kill container",getusernamerr)
		  return errors.ErrInternalserver
	}

	killerr := docker.Killcontainer(*killinstance,username)

	if killerr != nil {
		 return errors.ErrInternalserver
	}

	delerr := postgres.Containerkill(username,*killinstance)

	if delerr != nil {
		return errors.ErrInstancenotfound
	}

	return nil 
}