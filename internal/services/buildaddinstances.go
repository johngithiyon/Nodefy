package services

import (

	"github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/models"
	"github.com/johngithiyon/Nodefy/internal/repository/docker"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/postgres"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/redis"
)

func Addinstanceservices(sessionid string,Addinstances *models.Addinstancesbuild) error {

	//get username from redis 

	username,geterr := redis.Getusername(sessionid)

	if geterr != nil {
		return errors.ErrInternalserver
	}

	addinstancerr := docker.Addinstances(username,*Addinstances)

    if addinstancerr != nil {
		 return  errors.ErrInternalserver
	}

	dberr := postgres.SaveInstance(Addinstances.Instancename,username)

    if dberr != nil {
		return errors.ErrInternalserver
	}

	return  nil 
}