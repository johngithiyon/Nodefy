package services

import (
	"log"

	"github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/models"
	"github.com/johngithiyon/Nodefy/internal/repository/docker"
	"github.com/johngithiyon/Nodefy/internal/repository/linux"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/redis"
)

func DeployInstances(sessionid string,Deploy *models.Deploy ) error  {

   //get username from redis 

   username,geterr := redis.Getusername(sessionid)

	if geterr != nil {
		return errors.ErrInternalserver
	}

    // validate the user request

	if Deploy.OsName != "ubuntu" && Deploy.OsName != "alphine" {
			
		log.Println("Invalid os")
		return errors.ErrBadrequest
	}

	// validate the user selected services

	for _,services := range Deploy.Services {
	
		if services != "" && services != "postgresql" && services != "redis-server" {

			log.Println("Invalid Services")
			return errors.ErrBadrequest
		}
	}


   builerr :=  docker.BuildImage(username,Deploy)

   if builerr != nil {
	     log.Println("Error in the build image",builerr)
	     return errors.ErrInternalserver
   }

   userexists := linux.CheckUserlinux(username)

   if userexists != nil {

   adduserr := linux.Adduser(username)

   if adduserr != nil {
	   return adduserr
   }

	return nil 
  }

  return nil 
}