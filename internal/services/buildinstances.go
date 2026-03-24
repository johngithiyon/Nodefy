package services

import (
	"github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/models"
	"github.com/johngithiyon/Nodefy/internal/repository/docker"
	"github.com/johngithiyon/Nodefy/internal/repository/linux"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/redis"
	"github.com/johngithiyon/Nodefy/pkg/utils"
)

func BuildInstances(sessionid string,Build *models.Build ) error  {

   //get username from redis 

   username,geterr := redis.Getusername(sessionid)

	if geterr != nil {
		return errors.ErrInternalserver
	}

    //create a  network

	lowercase_username := utils.Lowercase(username)

	checkerr := docker.Checknetwork(username)

	if checkerr != nil { 

	createrr := docker.Createnetwork(lowercase_username)

	if createrr != nil {
		return  errors.ErrInternalserver
	}
  }	

	services := Build.Services

   for i:=0;i<len(services);i++ {
	      docker.BuildImage(username,services[i])
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