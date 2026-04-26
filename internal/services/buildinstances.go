package services

import (
	"encoding/json"
	"log"

	"github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/models"
	"github.com/johngithiyon/Nodefy/internal/repository/docker"
	"github.com/johngithiyon/Nodefy/internal/repository/linux"
	"github.com/johngithiyon/Nodefy/internal/repository/rabbitmq"
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


	payload := map[string]interface{}{
		"instancename":Build.Instancename,
		"services":Build.Services,
		"username":username,
	}

	data,converr := json.Marshal(payload)

	if converr != nil {
	return converr
	}

	chl,chlerr :=  rabbitmq.Createchannel()

	if chlerr != nil {
	  log.Println("channel err",chlerr)
	  return chlerr
	}

	publerr := rabbitmq.Publish(chl,"build_queue",data)

	if publerr != nil {
	   return  publerr
	}

	consumertype := "buildinstance"

	consumerr := rabbitmq.Consumer(data,consumertype)

	if consumerr != nil {
		log.Println("Err comes in",consumerr)
		return consumerr
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