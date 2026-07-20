package services

import (
	"github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/models"
	"github.com/johngithiyon/Nodefy/internal/repository/docker"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/postgres"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/redis"
)

func Exposeapp(sessionid string,Deployform *models.Exposeappform) error {

		username,geterr := redis.Getusername(sessionid)

		if geterr != nil {
			return errors.ErrInternalserver
		} 

        containerip,findiperr := docker.Findcontainerip(Deployform.Appname)

		if findiperr != nil {
			 return findiperr
		}

		Deployform.Username = username

		Deployform.Containerip = containerip

       saverr :=  postgres.ExposeDeployinstances(Deployform)

	   if saverr != nil {
		    return saverr
	   }

	    return  nil 
}