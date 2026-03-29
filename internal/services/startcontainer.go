package services

import (
	"github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/models"
	"github.com/johngithiyon/Nodefy/internal/repository/docker"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/postgres"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/redis"
)

func Startcontainerserives(sessionid string,startcontainer *models.Containermanage) error {

	//get username from redis 

	username,geterr := redis.Getusername(sessionid)

	if geterr != nil {
		return errors.ErrInternalserver
	}

	starterr := docker.Startcontainer(username,*startcontainer)

	if starterr != nil {
		return errors.ErrInternalserver
	}

    dberr := postgres.Containerstatus(username,*startcontainer)
	
	if dberr != nil {
		return errors.ErrInternalserver
	}

	return  nil 
}