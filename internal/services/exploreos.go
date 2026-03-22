package services

import (
	"github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/models"
	"github.com/johngithiyon/Nodefy/internal/repository/docker"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/redis"
	"github.com/johngithiyon/Nodefy/pkg/utils"
)

func Exploreos(sessionid string,exploreos *models.Exploreos) error {

	   
   //get username from redis 

   username,geterr := redis.Getusername(sessionid)

   if geterr != nil {
	   return errors.ErrInternalserver
   }

   //create a  network

   lowercase_username := utils.Lowercase(username)

   createrr := docker.Createnetwork(lowercase_username)

   if createrr != nil {
	   return  errors.ErrInternalserver
   }

  explorerr :=  docker.Exploreos(lowercase_username,*exploreos)

  if explorerr != nil {
	 return  errors.ErrInternalserver
  }
	 
	   return  nil 
}