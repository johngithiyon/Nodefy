package services

import (
	"github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/models"
	"github.com/johngithiyon/Nodefy/internal/repository/docker"
	"github.com/johngithiyon/Nodefy/internal/repository/linux"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/postgres"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/redis"
)

func Deployservices(sessionid string,Deploy *models.Deploy) error {
     
   //get username from redis 

   username,geterr := redis.Getusername(sessionid)

   if geterr != nil {
	   return errors.ErrInternalserver
   }

 
   deployerr :=  docker.Deploydocker(username,*Deploy)

   if deployerr != nil {
	    return  errors.ErrInternalserver
   }

   saverr := postgres.SaveDeployinstances(username,*Deploy)

   if saverr != nil {
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
   
	 return  nil 
}