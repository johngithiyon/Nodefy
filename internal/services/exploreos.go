package services

import (
	"encoding/json"
	"log"

	"github.com/johngithiyon/Nodefy/internal/errors"
	"github.com/johngithiyon/Nodefy/internal/models"
	"github.com/johngithiyon/Nodefy/internal/repository/linux"
	"github.com/johngithiyon/Nodefy/internal/repository/rabbitmq"
	"github.com/johngithiyon/Nodefy/internal/repository/storage/redis"
)

func Exploreos(sessionid string,exploreos *models.Exploreos) error {

	   
   //get username from redis 

   username,geterr := redis.Getusername(sessionid)

   if geterr != nil {
	   return errors.ErrInternalserver
   }

   payload := map[string]interface{}{
            "osname":exploreos.Osname,
            "instancename":exploreos.Instancename,
            "username":exploreos.Username,
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

 publerr := rabbitmq.Publish(chl,"explore_queue",data)

 if publerr != nil {
   return  publerr
 }

 consumertype := "exploreos"

  rabbitmq.Consumer(data,consumertype)

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