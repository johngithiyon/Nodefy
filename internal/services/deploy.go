package services

import (
	"encoding/json"
	"log"

	"github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/models"
	"github.com/johngithiyon/Nodefy/internal/repository/linux"
	"github.com/johngithiyon/Nodefy/internal/repository/rabbitmq"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/postgres"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/redis"
)

func Deployservices(sessionid string,Deploy *models.Deploy) error {
     
   //get username from redis 

   username,geterr := redis.Getusername(sessionid)

   if geterr != nil {
	   return errors.ErrInternalserver
   }


	payload := map[string]interface{}{
	  "username":username,
      "appname":Deploy.Appname,
      "gitrepo":Deploy.Gitrepo,
      "languages":Deploy.Languages,
      "services":Deploy.Services,
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

	publerr := rabbitmq.Publish(chl,"deploy_queue",data)

	if publerr != nil {
	   return  publerr
	}

	consumertype := "deployinstance"

	consumerr := rabbitmq.Consumer(data,consumertype)

	if consumerr != nil {
		log.Println("Err comes in",consumerr)
		return consumerr
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