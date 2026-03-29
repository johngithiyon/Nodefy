package services

import (
	"log"

	"github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/models"
	"github.com/johngithiyon/Nodefy/internal/repository/docker"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/postgres"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/redis"
)

func Stopcontainerservices(sessionid string,Stopcontainer *models.Containermanage) error{

	username,getusernamerr := redis.Getusername(sessionid)

	if username == "" && getusernamerr != nil {
		  log.Println("Get username err from the stop container",getusernamerr)
		  return errors.ErrInternalserver
	}

	  dockerstoperr :=   docker.Stopcontainer(username,*Stopcontainer)

	  if dockerstoperr != nil {
		    return errors.ErrInternalserver
	  }

	  dberr :=  postgres.Containerstatus(username,*Stopcontainer)

	  if dberr != nil {
		    return errors.ErrInternalserver
	  }

	  return  nil 
}