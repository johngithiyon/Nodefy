package services

import (
	"github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/models"
	"github.com/johngithiyon/Nodefy/internal/repository/docker"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/postgres"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/redis"
)

func DeployAddinstanceservices(sessionid string,DeployAddinstances *models.Deployaddinstances) error {

	//get username from redis 

	username,geterr := redis.Getusername(sessionid)

	if geterr != nil {
		return errors.ErrInternalserver
	}

	addinstancerr := docker.DeployAddinstances(username,*DeployAddinstances)

    if addinstancerr != nil {
		 return  errors.ErrInternalserver
	}

	saverr := postgres.SaveDeployAddinstances(username,*DeployAddinstances)

	if saverr != nil {
		return  errors.ErrInternalserver
	}
 
	return  nil 

}