package services

import (
	"github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/models"
	"github.com/johngithiyon/Nodefy/internal/repository/docker"
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

   saverr := postgres.Savedeployinstance(username,*Deploy)

   if saverr != nil {
	  return errors.ErrInternalserver
   }
   
	 return  nil 
}