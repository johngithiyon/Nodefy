package redis

import (
	"context"
	"log"
	"time"
)

func Storesessionid(id string,username string ) error {
         
  	seterr := Redisconn.Redisconn.Set(context.Background(),id,username, 7*24*time.Hour).Err()

	if seterr != nil {
        log.Println("Set Err in session",seterr)
		return seterr 
	}

	return  nil 
}

func Getusername(id string) (string,error){

	 username,geterr :=  Redisconn.Redisconn.Get(context.Background(),id).Result()

	 if geterr != nil {
		   log.Println("Username Get Err",geterr)
		   return "",geterr
	 }

	 return username,nil
}